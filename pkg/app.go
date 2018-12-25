package pkg

import (
	"log"
	"net/http"
	"net/http/httptest"
)

// App struct is the central structure to manage the application
type App struct {
	conf   *Conf
	server *server
}

// Initialize is preparing the App struct
func (a *App) Initialize(c Conf) {
	a.conf = &c
	a.server = &server{}
	a.server.initialize(c)
}

// LogStartup is logging out at startup
func (a *App) LogStartup() {
	// startup logs
	log.Printf("Starting up snoop API")
	a.conf.PrintConf()
}

// Listen is starting to listen on the specified port
func (a *App) Listen() {
	a.server.listen()
}

// ExecuteTestRequest is executing the requests against the application
func (a *App) ExecuteTestRequest(req *http.Request) *httptest.ResponseRecorder {
	return a.server.executeTestRequest(req)
}
