package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request){
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintln(writer, "hello")
	} else {
		fmt.Fprintf(writer, "hello %s", name)
	}
}

func TestQueryParams(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/hello?name=haikal",nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)

}

func MultipleQueryParams(writer http.ResponseWriter, request *http.Request){
	fristname := request.URL.Query().Get("first_name")
	lastname := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "hello %s %s", fristname, lastname)

}

func TestMultipleQueryParams(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/hello?first_name=haikal&last_name=frastiawan",nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParams(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)

}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request){
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/hello?name=haikal&name=frastiawan&name=setiawan",nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)

}