package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = setupRouter()

func TestPingEndpoint(t *testing.T) {
	w := httptest.NewRecorder()
	resp := struct {
		Message string
	}{}

	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resp.Message, "pong")
}

func TestHelloEndpoint(t *testing.T) {
	w := httptest.NewRecorder()
	resp := struct {
		Message string
	}{}

	req, _ := http.NewRequest("GET", "/hello/test", nil)
	router.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resp.Message, "Hello test")
}
