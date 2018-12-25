package main

import (
	"net/http"
	"testing"

	"github.com/bboortz/snoop-api/pkg"
	"github.com/stretchr/testify/assert"
)

var a pkg.App

func TestMain(t *testing.M) {

	c := pkg.Conf{
		Port:         ":8888",
		Protocol:     "http",
		ReadTimeout:  5,
		WriteTimeout: 10,
	}

	a = pkg.App{}
	a.Initialize(c)

	t.Run()
}

func TestAppUp(t *testing.T) {
	ass := assert.New(t)

	req, _ := http.NewRequest("GET", "/", nil)
	resp := a.ExecuteTestRequest(req)

	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
}
