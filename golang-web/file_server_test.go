package golangweb

import (
	"embed"
	"net/http"
	"io/fs"
	"testing"
)

func TestFileServer(t *testing.T) {

	directory := http.Dir("./resource")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/",  http.StripPrefix("/static", fileServer))

	server := http.Server {
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resource
var resource embed.FS

func TestFileServerGolangEmbed(t *testing.T) {

	directory, _ := fs.Sub(resource, "resource")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/",  http.StripPrefix("/static", fileServer))

	server := http.Server {
		Addr: "localhost:8080",
		Handler: mux,
	}
	defer server.Close()

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}