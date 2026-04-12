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

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive Request")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Request")
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func (w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
		fmt.Fprint(w, "Hello Middleware")
	})

	middleware := LogMiddleware{Handler: router}


	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)


	assert.Equal(t, "Hello Middleware", string(body))

}