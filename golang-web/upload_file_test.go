package golangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	// request.ParseForm(100 << 20)
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := request.FormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUpload(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, _ := writer.CreateFormFile("file", "sample-product.png.jpg")
	file.Write([]byte("Sample image content"))
	writer.WriteField("name", "Haikal")
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	response := recorder.Result()
	bodyResponse, _ := io.ReadAll(response.Body)
	fmt.Println(string(bodyResponse))
}

// TestUploadForm starts a real server for manual testing.
// Note: This will block and cause 'go test' to stay running.
func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	fmt.Println("Server started at http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

