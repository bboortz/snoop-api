package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDumpNil(t *testing.T) {
	ass := assert.New(t)

	requestDump := requestDump{}
	requestDump.initialze(nil)

	// check requestDump
	ass.Equalf("", requestDump.Method, "Invalid Method")
	ass.Equalf("", requestDump.Protocol, "Invalid Protocol")
	ass.Equalf("", requestDump.Path, "Invalid Path")
}

func TestDumpSimple(t *testing.T) {
	ass := assert.New(t)

	// prepare requestDump
	req, _ := http.NewRequest("GET", "/", nil)
	requestDump := requestDump{}
	requestDump.initialze(req)

	// check requestDump
	ass.Equalf("GET", requestDump.Method, "Invalid Method")
	ass.Equalf("HTTP/1.1", requestDump.Protocol, "Invalid Protocol")
	ass.Equalf("/", requestDump.Path, "Invalid Path")
}

func TestDumpComplex(t *testing.T) {
	ass := assert.New(t)

	// preparing te request
	type User struct {
		Name string
	}
	user := &User{Name: "Frank"}
	b, _ := json.Marshal(user)

	// prepare requestDump
	req, _ := http.NewRequest("GET", "/complex?key1=value1&key2=value2", bytes.NewBuffer(b))
	req.Header.Set("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	requestDump := requestDump{}
	requestDump.initialze(req)

	// check requestDump
	ass.Equalf("GET", requestDump.Method, "Invalid Method")
	ass.Equalf("HTTP/1.1", requestDump.Protocol, "Invalid Protocol")
	ass.Equalf("/complex", requestDump.Path, "Invalid Path")
	ass.Equalf(2, len(requestDump.Header), "Invalid Number of Header")
	ass.Equalf("application/json", requestDump.Header["Accept"], "Invalid Header \"Accept\"")
	ass.Equalf(2, len(requestDump.Parameter), "Invalid Number of Parameter")
	ass.Equalf("value1", requestDump.Parameter["key1"], "Invalid Parameter \"key1\"")
	ass.Equalf("value2", requestDump.Parameter["key2"], "Invalid Parameter \"key2\"")
}
