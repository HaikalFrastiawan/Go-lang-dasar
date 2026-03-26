package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	fristname := r.PostForm.Get("first_name")
	lastname := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s",fristname, lastname )
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Haikal&last_name=Frastiawan")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)


}