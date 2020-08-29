package api_test

import (
	"net/http"
	"testing"
)

func TestGetMessages(t *testing.T) {
	tests := []httpTest{
		{
			testName:       "[POST] Normal",
			body:           ``,
			expectedStatus: http.StatusOK,
			url:            "/api/messages/queue1",
			expectedBody:   testData("messages_normal.out"),
		},
		{
			testName:       "[GET] Normal",
			body:           ``,
			expectedStatus: http.StatusOK,
			url:            `/api/messages/?queue=queue1`,
			expectedBody:   testData("messages_normal.out"),
		},
		{
			testName:       "[POST] Queue not found",
			body:           ``,
			expectedStatus: http.StatusBadRequest,
			url:            "/api/messages/q",
			expectedBody:   "Queue 'q' not found",
		},
	}

	withServer(t, tests, "messages", func(t *testing.T, test httpTest) {
		if err := test.run(); err != nil {
			t.Fatal(err)
		}
	})
}
