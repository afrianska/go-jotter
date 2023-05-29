package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/afrianska/playgo/unit-test/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	tests := []struct {
		description string
		route string
		expectedError bool
		expectedCode int
		expextedBody string
	}{
		{
			description: "index route",
			route: "/",
			expectedError: false,
			expectedCode: 200,
			expextedBody: "Welcome, You are at home page! dsdssfasd",
		},
		{
			description: "non existing route",
			route: "/i-dont-exist",
			expectedError: false,
			expectedCode: 404,
			expextedBody: "Cannot Get /i-dont-exitst",
		},
	}

	app := utils.Setup()

	for _, test := range tests {
		req, _ := http.NewRequest(
			"Get",
			test.route,
			nil,
		)

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		body, err := io.ReadAll(res.Body)

		assert.Nilf(t, err, test.description)

		assert.Equalf(t, test.expextedBody, string(body), test.description)
	}
}