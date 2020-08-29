package api_test

import (
	"net/http"
	"testing"
)

func TestPublish(t *testing.T) {
	tests := []httpTest{
		{
			testName: "[POST] Normal",
			body: `{
				"message": "publishing message"
			}`,
			expectedStatus: http.StatusOK,
			url:            "/api/publish/queue1",
			expectedBody:   "",
		},
		{
			testName:       "[GET] Normal",
			expectedStatus: http.StatusOK,
			url:            "/api/publish/?queue=queue1&msg=anymessage",
			expectedBody:   "",
		},
	}

	withServer(t, tests, "messages", func(t *testing.T, test httpTest) {
		if err := test.run(); err != nil {
			t.Fatal(err)
		}
	})
}
