package pkg

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	Router    *mux.Router
	conf      *Conf
	tlsConfig *tls.Config
}

func (s *server) initialize(c Conf) {
	s.conf = &c
	s.Router = mux.NewRouter()
	s.setupRoutes()
	s.setupMiddleware()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log requests
		log.Printf("Processing request: %v %v %v", r.Method, r.URL.Path, r.Proto)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func (s *server) setupRoutes() {
	// catch all requests
	s.Router.PathPrefix("/health").HandlerFunc(health)
	//s.Router.PathPrefix("/").HeadersRegexp("accept", "application/json").HandlerFunc(indexJSON)
	s.Router.PathPrefix("/").HandlerFunc(indexJSON)
}

func (s *server) setupMiddleware() {
	s.Router.Use(loggingMiddleware)
}

func (s *server) setupTLS() {
	// for the list of TLS settings please refer to https://golang.org/pkg/crypto/tls/
	tlsCipherMap := map[string]uint16{
		"TLS_RSA_WITH_RC4_128_SHA":                tls.TLS_RSA_WITH_RC4_128_SHA,
		"TLS_RSA_WITH_3DES_EDE_CBC_SHA":           tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		"TLS_RSA_WITH_AES_128_CBC_SHA":            tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		"TLS_RSA_WITH_AES_256_CBC_SHA":            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		"TLS_RSA_WITH_AES_128_CBC_SHA256":         tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
		"TLS_RSA_WITH_AES_128_GCM_SHA256":         tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		"TLS_RSA_WITH_AES_256_GCM_SHA384":         tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		"TLS_ECDHE_ECDSA_WITH_RC4_128_SHA":        tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
		"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA":    tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA":    tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		"TLS_ECDHE_RSA_WITH_RC4_128_SHA":          tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
		"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA":     tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
		"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA":      tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA":      tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256": tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
		"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256":   tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
		"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256":   tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256": tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384":   tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384": tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305":    tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305":  tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		"TLS_FALLBACK_SCSV":                       tls.TLS_FALLBACK_SCSV,
	}

	tlsVersionMap := map[string]uint16{
		"VersionSSL30": tls.VersionSSL30,
		"VersionTLS10": tls.VersionTLS10,
		"VersionTLS11": tls.VersionTLS11,
		"VersionTLS12": tls.VersionTLS12,
	}

	ciphers := strings.Split(s.conf.TLSCipher, "|")
	cipherSuites := []uint16{}
	for _, cipher := range ciphers {
		cipherSuites = append(cipherSuites, tlsCipherMap[cipher])
	}

	s.tlsConfig = &tls.Config{
		CipherSuites:             cipherSuites,
		PreferServerCipherSuites: true,
		MinVersion:               tlsVersionMap[s.conf.TLSMinVersion],
	}
	s.tlsConfig.BuildNameToCertificate()
}

func (s *server) listen() {
	// for timeout settings please refer to https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	httpServer := &http.Server{
		Handler:      s.Router,
		Addr:         s.conf.Port,
		ReadTimeout:  time.Duration(s.conf.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.conf.WriteTimeout) * time.Second,
	}

	if s.conf.Protocol == "https" {
		httpServer.TLSConfig = s.tlsConfig
		log.Fatal(httpServer.ListenAndServeTLS(s.conf.TLSCert, s.conf.TLSKey))
	} else {
		log.Fatal(httpServer.ListenAndServe())
	}

}

func (s *server) executeTestRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}
