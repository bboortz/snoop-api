package main

import (
	"github.com/bboortz/snoop-api/pkg"
)

func main() {
	// initialize config with default settings
	c := pkg.Conf{
		Port:         ":8443",
		Protocol:     "http",
		ReadTimeout:  5,
		WriteTimeout: 10,
	}

	// load the config
	c.LoadConf("snoop.yaml")

	// start listener
	a := pkg.App{}
	a.Initialize(c)
	a.LogStartup()
	a.Listen()
}
