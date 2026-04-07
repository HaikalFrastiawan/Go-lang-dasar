package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func (w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
		fmt.Fprint(w, "Hello http router")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	server.ListenAndServe()
}