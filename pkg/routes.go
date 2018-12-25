package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

type content struct {
	Request  *http.Request  `json:"request"`
	Response *http.Response `json:"response"`
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Printf("health")
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

// indexText is serving all text requests /*
func indexText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, string(requestDump))
	}
}

// indexJson is serving json requests /*
func indexJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dump := requestDump{}
	dump.initialze(r)
	json.NewEncoder(w).Encode(dump)
}
