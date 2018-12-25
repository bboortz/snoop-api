package pkg

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type requestDump struct {
	RemoteAddress    string            `json:"remoteAddr"`
	Method           string            `json:"method"`
	Protocol         string            `json:"protocol"`
	Path             string            `json:"path"`
	Host             string            `json:"host"`
	ContentLength    int64             `json:"contentLength"`
	TransferEncoding string            `json:"transferEncoding"`
	Connection       string            `json:"connection"`
	Header           map[string]string `json:"header"`
	Body             string            `json:"body"`
	Parameter        map[string]string `json:"parameter"`
}

func (d *requestDump) initialze(req *http.Request) {
	if req != nil {
		d.RemoteAddress = req.RemoteAddr
		d.Method = req.Method
		d.Protocol = req.Proto
		d.Path = req.URL.Path
		d.Host = req.Host
		d.ContentLength = req.ContentLength

		d.setConnection(req)
		d.setTransferEncoding(req)
		d.setHeader(req)
		d.setBody(req)
		d.setParameter(req)
	}
}

func (d *requestDump) setTransferEncoding(req *http.Request) {
	if len(req.TransferEncoding) > 0 {
		d.TransferEncoding = strings.Join(req.TransferEncoding, ",")
	}
}

func (d *requestDump) setConnection(req *http.Request) {
	if req.Close {
		d.Connection = "close"
	}
}

func (d *requestDump) setHeader(req *http.Request) {
	m := make(map[string]string)

	for key, value := range req.Header {
		m[key] = value[0]
	}
	d.Header = m
}

func (d *requestDump) setBody(req *http.Request) {
	if req.Body != nil {
		// defer req.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(req.Body)
		if len(bodyBytes) > 0 {
			bodyString := string(bodyBytes)
			d.Body = bodyString
		}
	}
}

func (d *requestDump) setParameter(req *http.Request) {
	m := make(map[string]string)
	keys := req.URL.Query()

	for key := range keys {
		m[key] = keys.Get(key)
	}
	d.Parameter = m
}
