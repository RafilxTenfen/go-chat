package api_test

import (
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	tests := []httpTest{
		{testName: "[POST] Normal",
			body:           `{"email": "adminxt@gmail.com", "password": "defaultpwd"}`,
			expectedStatus: http.StatusOK,
			url:            "/login",
			expectedBody:   testData("login.out"),
		},
		{testName: "[POST] Invalid Email",
			body:           `{"email": "adminxt13213@gmail.com", "password": "defaultpwd"}`,
			expectedStatus: http.StatusNotFound,
			url:            "/login",
			expectedBody:   "This email adminxt13213@gmail.com doesn't exists",
		},
	}

	withServer(t, tests, "users", func(t *testing.T, test httpTest) {
		if err := test.run(); err != nil {
			t.Fatal(err)
		}
	})
}
