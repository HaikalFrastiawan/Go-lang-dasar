package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotfound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not Found")
	})
	router.GET("/", func (w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
		fmt.Fprint(w, "Hello http router")
	})


	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/404", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)


	assert.Equal(t, "Not Found", string(body))

}