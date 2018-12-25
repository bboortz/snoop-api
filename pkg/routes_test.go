package pkg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteGetIndexText(t *testing.T) {
	ass := assert.New(t)

	// sending the request
	req, _ := http.NewRequest("GET", "/", nil)
	resp := executeRequest(req, indexText)

	// check response header
	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
	ass.Equalf("application/text", resp.HeaderMap["Content-Type"][0], "Invalid Content-Type")

	// check response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	requestDump := &requestDump{}
	err := json.Unmarshal([]byte(bodyString), requestDump)
	ass.NotNilf(err, "No error occured during unmarshalling of normal text")
}

func TestRouteGetIndexJSON(t *testing.T) {
	ass := assert.New(t)

	// sending the request
	req, _ := http.NewRequest("GET", "/", nil)
	resp := executeRequest(req, indexJSON)

	// check response header
	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
	ass.Equalf("application/json", resp.HeaderMap["Content-Type"][0], "Invalid Content-Type")

	// check response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	requestDump := &requestDump{}
	_ = json.Unmarshal([]byte(bodyString), requestDump)

	ass.Equalf("HTTP/1.1", requestDump.Protocol, "Invalid Protocol")
	ass.Equalf("/", requestDump.Path, "Invalid Path")
	ass.Equalf(0, len(requestDump.Header), "Invalid Number of Header")
	ass.Equalf(0, len(requestDump.Parameter), "Invalid Number of Parameter")
}

func TestRouteGetIndexJSONComplex(t *testing.T) {
	ass := assert.New(t)

	// sending the request
	req, _ := http.NewRequest("GET", "/complex?key1=value1&key2=value2", nil)
	req.Header.Set("Accept", "application/json")
	resp := executeRequest(req, indexJSON)

	// check response header
	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
	ass.Equalf("application/json", resp.HeaderMap["Content-Type"][0], "Invalid Content-Type")

	// check response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	requestDump := &requestDump{}
	_ = json.Unmarshal([]byte(bodyString), requestDump)

	ass.Equalf("HTTP/1.1", requestDump.Protocol, "Invalid Protocol")
	ass.Equalf("/complex", requestDump.Path, "Invalid Path")
	ass.Equalf(1, len(requestDump.Header), "Invalid Number of Header")
	ass.Equalf("application/json", requestDump.Header["Accept"], "Invalid Header \"Accept\"")
	ass.Equalf(2, len(requestDump.Parameter), "Invalid Number of Parameter")
	ass.Equalf("value1", requestDump.Parameter["key1"], "Invalid Parameter \"key1\"")
	ass.Equalf("value2", requestDump.Parameter["key2"], "Invalid Parameter \"key2\"")
}

func TestRoutePostIndexJSONComplex(t *testing.T) {
	ass := assert.New(t)

	// preparing te request
	type User struct {
		Name string
	}
	user := &User{Name: "Frank"}
	b, _ := json.Marshal(user)

	// sending the request
	req, _ := http.NewRequest("GET", "/complex", bytes.NewBuffer(b))
	req.Header.Set("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp := executeRequest(req, indexJSON)

	// check response header
	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
	ass.Equalf("application/json", resp.HeaderMap["Content-Type"][0], "Invalid Content-Type")

	// check response body
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	requestDump := &requestDump{}
	_ = json.Unmarshal([]byte(bodyString), requestDump)

	ass.Equalf("HTTP/1.1", requestDump.Protocol, "Invalid Protocol")
	ass.Equalf("/complex", requestDump.Path, "Invalid Path")
	ass.Equalf(2, len(requestDump.Header), "Invalid Number of Header")
	ass.Equalf(0, len(requestDump.Parameter), "Invalid Number of Parameter")
	ass.Equalf(16, len(requestDump.Body), "Invalid Number Body")
}

func executeRequest(req *http.Request, handlerFunc http.HandlerFunc) *httptest.ResponseRecorder {
	// Creating the ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	return rr
}
