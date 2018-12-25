package pkg

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var a App

func TestMain(m *testing.M) {

	c := Conf{
		Port:          ":8444",
		Protocol:      "https",
		ReadTimeout:   5,
		WriteTimeout:  10,
		TLSCipher:     "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256|TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
		TLSMinVersion: "VersionTLS12",
		TLSCert:       "./examples/certs/cert.pem",
		TLSKey:        "./examples/certs/key.pem",
	}

	a = App{}
	a.Initialize(c)
	a.LogStartup()

	os.Exit(m.Run())
}

func TestAppUp(t *testing.T) {
	ass := assert.New(t)

	req, _ := http.NewRequest("GET", "/", nil)
	resp := a.ExecuteTestRequest(req)

	ass.Equalf(http.StatusOK, resp.Code, "Invalid Response Code")
}
