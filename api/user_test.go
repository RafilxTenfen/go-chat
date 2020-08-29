package api_test

import (
	"net/http"
	"testing"
)

func TestGetAllUser(t *testing.T) {
	tests := []httpTest{
		{testName: "[GET] Normal", url: "/api/user", expectedBody: testData("get_all_users_normal.out")},
	}

	withServer(t, tests, "users", func(t *testing.T, test httpTest) {
		if err := test.run(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestGetUser(t *testing.T) {
	tests := []httpTest{
		{testName: "[GET] Normal", url: "/api/user/6tXaLO1p4I8cLbGkgl8Jgy", expectedBody: testData("get_users_normal.out")},
	}

	withServer(t, tests, "users", func(t *testing.T, test httpTest) {
		if err := test.run(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestInsertUser(t *testing.T) {
	tests := []httpTest{
		{testName: "[POST] Normal",
			body: `{
				"email": "adminx@gmail.com", 
				"password": "defaultpwd1"
			}`,
			expectedStatus: http.StatusOK,
			expectedBody:   testData("insert_user_normal.out"),
		},
		{testName: "[POST] Invalid PWD",
			body: `{
				"email": "adminx@gmail.com" 
			}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Password should be at least 6 characteres",
		},
	}

	withServer(t, tests, "users", func(t *testing.T, test httpTest) {
		test.url = "/api/user"
		if err := test.run(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	tests := []httpTest{
		{
			testName:       "[DELETE] Normal",
			url:            "/api/user/%s",
			urlParams:      []interface{}{"2cWxYz0I4jae3KFuuNpuhy"},
			expectedStatus: http.StatusOK,
		},
		{
			testName:       "[DELETE] Not exists",
			url:            "/api/user/%s",
			urlParams:      []interface{}{"NonExistEnt1337"},
			expectedStatus: http.StatusOK,
		},
	}

	withServer(t, tests, "users", func(t *testing.T, test httpTest) {
		if err := test.run(); err != nil {
			t.Fatal(err)
		}
	})
}

func TestUpdateUser(t *testing.T) {
	tests := []httpTest{
		{
			testName:       "[PUT] Normal",
			body:           `{"email":"adminx20112@agmail.com"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   testData("update_user_normal.out"),
		},
	}

	withServer(t, tests, "users", func(t *testing.T, test httpTest) {
		test.url = "/api/user/6tXaLO1p4I8cLbGkgl8Jgy"
		if err := test.run(); err != nil {
			t.Fatal(err)
		}
	})
}
