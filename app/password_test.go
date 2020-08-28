package app_test

import (
	"fmt"
	"testing"

	"github.com/RafilxTenfen/go-chat/app"
)

func TestGeneratePwd(t *testing.T) {
	tests := []struct {
		email       string
		pwd         string
		expectedErr error
	}{
		{
			email:       "mybaseemail@gmail.com",
			pwd:         "mydummypwd",
			expectedErr: nil,
		},
	}

	for i, test := range tests {
		_, receivedErr := app.GeneratePwd(test.email, test.pwd)
		if receivedErr != test.expectedErr {
			t.Errorf("Error on test %d \nExpectedErr: %+v \nReceived Err: %+v", i, test.expectedErr, receivedErr)
		}
	}
}

func TestComparePasswordAndHash(t *testing.T) {
	tests := []struct {
		email         string
		pwd           string
		hash          string
		expectedMatch bool
		expectedErr   error
	}{
		{
			email:         "mybaseemail@gmail.com",
			pwd:           "mydummypwd",
			hash:          "$argon2id$v=19$m=65536,t=1,p=2$lP+K517qHqsPduGTT/Y4ow$tXKS+hLGeHQOgNhpPSeTZdZ5hz811MSZmmAgLy9ebzU",
			expectedMatch: true,
			expectedErr:   nil,
		},
		{
			email:         "mybaseemail@gmail.com",
			pwd:           "mydummypwd1",
			hash:          "$argon2id$v=19$m=65WrongHash",
			expectedMatch: false,
			expectedErr:   fmt.Errorf("argon2id: hash is not in the correct format"),
		},
	}

	for i, test := range tests {
		receivedMatch, receivedErr := app.ComparePasswordAndHash(test.email, test.pwd, test.hash)
		if receivedMatch != test.expectedMatch || (receivedErr != test.expectedErr && receivedErr.Error() != test.expectedErr.Error()) {
			t.Errorf("Error on test %d \nExpectedMatch: %t \nReceivedMatch: %t \nExpectedErr: %+v \nReceivedErr: %+v", i, test.expectedMatch, receivedMatch, test.expectedErr, receivedErr)
		}
	}
}
