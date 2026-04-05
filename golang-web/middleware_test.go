package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before request")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After request")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (ErrorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()
	ErrorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Middleware")
	})
	logMiddleware := &LogMiddleware{Handler: mux}

	ErrorHandler := &ErrorHandler{Handler: logMiddleware}

	
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: ErrorHandler,
	}
	err := server.ListenAndServe()
	if err != nil {	
		panic(err)
	}	
}