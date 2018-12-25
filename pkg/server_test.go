package pkg

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerUp(t *testing.T) {
	ass := assert.New(t)

	req, _ := http.NewRequest("GET", "/", nil)
	resp := a.server.executeTestRequest(req)

	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
}

func TestServerGetIndexText(t *testing.T) {
	ass := assert.New(t)

	req, _ := http.NewRequest("GET", "/", nil)
	resp := a.server.executeTestRequest(req)

	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
	ass.Equalf("application/json", resp.HeaderMap["Content-Type"][0], "Invalid Content-Type")
}

func TestServerGetIndexJSON(t *testing.T) {
	ass := assert.New(t)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json")
	resp := a.server.executeTestRequest(req)

	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
	ass.Equalf("application/json", resp.HeaderMap["Content-Type"][0], "Invalid Content-Type")
}

func TestServerPostIndexText(t *testing.T) {
	ass := assert.New(t)

	req, _ := http.NewRequest("POST", "/", nil)
	resp := a.server.executeTestRequest(req)

	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
	ass.Equalf("application/json", resp.HeaderMap["Content-Type"][0], "Invalid Content-Type")
}

func TestServerPOSTIndexJSON(t *testing.T) {
	ass := assert.New(t)

	req, _ := http.NewRequest("POST", "/", nil)
	req.Header.Set("Accept", "application/json")
	resp := a.server.executeTestRequest(req)

	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
	ass.Equalf("application/json", resp.HeaderMap["Content-Type"][0], "Invalid Content-Type")
}

func TestServerIsCatchingAllRequests(t *testing.T) {
	ass := assert.New(t)

	req, _ := http.NewRequest("GET", "/anotherpath", nil)
	resp := a.server.executeTestRequest(req)

	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
}
