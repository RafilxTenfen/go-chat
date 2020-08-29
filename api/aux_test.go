package api_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/NeowayLabs/wabbit/amqptest/server"
	"github.com/RafilxTenfen/go-chat/api"
	"github.com/RafilxTenfen/go-chat/chat"
	"github.com/RafilxTenfen/go-chat/internal/dbtest"
	"github.com/RafilxTenfen/go-chat/rabbit"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	echoLog "github.com/labstack/gommon/log"
	"github.com/rhizomplatform/fs"
	"github.com/rhizomplatform/log"
)

type httpTest struct {
	testName  string
	method    string
	url       string
	urlParams []interface{}
	body      string
	form      string

	expectedStatus int
	expectedBody   interface{}

	actualBody string

	server *api.Server
	db     *gorm.DB
}

func withServer(t *testing.T, tests []httpTest, dep string, fn func(t *testing.T, test httpTest)) {
	baseFolder, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}

	log.Setup(fs.Path(baseFolder), "tests", 2, 1)
	log.SetStdoutLevel(log.LevelOff)
	log.SetFileLevel(log.LevelOff)

	dbtest.WithDB(func(db *gorm.DB) {
		dbtest.Require(db, dep)

		e := echo.New()
		e.Logger = api.LogrusAdapter()
		e.Logger.SetLevel(echoLog.ERROR)
		e.Use(api.Logger())

		amqpURL := "amqp://localhost:5672/%2f"
		fakeServer := server.NewServer(amqpURL)
		if err := fakeServer.Start(); err != nil {
			panic(err)
		}

		conn, ch, err := rabbit.Init(amqpURL)
		if err != nil {
			panic(err)
		}
		st := rabbit.Settings{
			QuantityMessageQueue: 2,
			RabbitMqURL:          "",
		}

		chatUsr := chat.NewUserChatStructure(nil, db, conn, ch, st)

		defer func() {
			if err := fakeServer.Stop(); err != nil {
				panic(err)
			}
		}()
		s := api.NewServer(e, db, chatUsr)

		// FIXME: maybe when roles are properly check,
		// the skip won't be necessary, the mock will suffice
		s.SetCustomContextKey("auth", false)
		s.MockJWTToken()

		s.Routes()

		for i := range tests {
			test := tests[i]
			test.server = s
			test.db = db

			if test.testName == "" {
				test.testName = fmt.Sprintf("Unnamed #%d", i)
			}

			t.Run(test.testName, func(t *testing.T) {
				fn(t, test)
			})
		}
	})

	log.TearDown()
	fs.RemoveAll(baseFolder)
}

func (tc *httpTest) run() error {
	return tc.runWith(tc.server)
}

func (tc *httpTest) runWith(s *api.Server) error {
	method := tc.method
	if method == "" {
		switch {
		case strings.Contains(tc.testName, "[POST]"):
			method = http.MethodPost
		case strings.Contains(tc.testName, "[PUT]"):
			method = http.MethodPut
		case strings.Contains(tc.testName, "[DELETE]"):
			method = http.MethodDelete
		default:
			method = http.MethodGet
		}
	}

	addr := tc.url
	if len(tc.urlParams) > 0 {
		addr = fmt.Sprintf(addr, tc.urlParams...)
	}

	if splitted := strings.Split(addr, "?"); len(splitted) > 1 {
		qs, err := url.ParseQuery(splitted[1])
		if err != nil {
			return err
		}

		addr = splitted[0] + "?" + qs.Encode()
	}

	if tc.expectedStatus == 0 {
		tc.expectedStatus = http.StatusOK
	}

	var payload io.Reader
	if tc.body != "" {
		payload = strings.NewReader(tc.body)
	}
	if tc.form != "" {
		payload = strings.NewReader(tc.form)
	}

	req := httptest.NewRequest(method, addr, payload)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	// check the http code
	if rec.Code != tc.expectedStatus {
		return fmt.Errorf("expected status code of %d, but received %d\nBody: %s", tc.expectedStatus, rec.Code, rec.Body.String())
	}

	// extract the body	content
	var expected string

	// ignore nil pointer
	v := reflect.ValueOf(tc.expectedBody)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}

		v = reflect.Indirect(v)
	}

	// ignore invalid value
	if !v.IsValid() {
		return nil
	}

	// if it is a struct, marshal it
	if v.Kind() == reflect.Struct {
		switch obj := v.Interface().(type) {
		case json.Marshaler:
			b, err := obj.MarshalJSON()
			if err != nil {
				return err
			}
			expected = string(b)

		default:
			b, err := json.Marshal(obj)
			if err != nil {
				return err
			}
			expected = string(b)
		}
	} else {
		switch obj := v.Interface().(type) {
		case string:
			expected = obj
		default:
			return fmt.Errorf("unknown type: %v", obj)
		}
	}

	// get the expected and the actual body
	expected = strings.TrimSpace(expected)
	body := strings.TrimSpace(rec.Body.String())

	// sometimes (mainly on inserts) the ID fields must be ignored
	handleID("id", &expected, body)
	handleID("uuid", &expected, body)
	handleID("token", &expected, body)
	handleID("createdAt", &expected, body)

	tc.actualBody = body

	if body != expected {
		return fmt.Errorf("expected body to be '%s', \n but it was '%s'", expected, body)
	}

	return nil
}

func handleID(fieldName string, expected *string, actual string) {
	field := `"` + fieldName + `":`

	if !strings.Contains(*expected, field+"?") {
		return
	}

	start := strings.Index(actual, field)
	if start == -1 {
		return
	}

	start += len(field)
	end := strings.IndexRune(actual[start+1:], '"')
	substr := actual[start : end+start+2]

	*expected = strings.Replace(*expected, field+"?", field+substr, 1)
}

func testData(filename string) string {
	f, err := os.Open("testdata/" + filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("error on close testdata file %+v", err)
		}
	}()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(b)
}
