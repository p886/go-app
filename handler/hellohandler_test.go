package handler

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHelloHandler(t *testing.T) {
	handler := NewHelloHandler()

	request := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, recorder.Result().Status, "200 OK")
	assert.Equal(t, responseBody, []byte("Hello World! This is an update"))
}
