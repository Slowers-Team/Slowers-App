package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlowersRoute(t *testing.T) {
	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "\"GET /api/flowers\" returns 200",
			route:         "/api/flowers",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
	}
}
