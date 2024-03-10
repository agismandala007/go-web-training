package web_1

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writter, "Hello World")
}

func TestHttp(t *testing.T) {
	newRequest := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, newRequest)

	respone := recorder.Result()
	body, _ := io.ReadAll(respone.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
