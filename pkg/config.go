package pkg

import (
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
)

// Conf is the struct to store the config settings
type Conf struct {
	Port          string `yaml:"port"`
	Protocol      string `yaml:"protocol"`
	ReadTimeout   int    `yaml:"readTimeout"`
	WriteTimeout  int    `yaml:"WriteTimeout"`
	TLSCipher     string `yaml:"TLSCipher"`
	TLSMinVersion string `yaml:"TLSMinVersion"`
	TLSCert       string `yaml:"TLSCert"`
	TLSKey        string `yaml:"TLSKey"`
}

// LoadConf is loading config settings from configFile
func (c *Conf) LoadConf(configFile string) *Conf {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("error loading config file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// PrintConf is printing config settings
func (c *Conf) PrintConf() {
	log.Printf("Port: \t\t%s\n", c.Port)
	log.Printf("Protocol: \t\t%s\n", c.Protocol)
	log.Printf("ReadTimeout: \t%d\n", c.ReadTimeout)
	log.Printf("WriteTimeout: \t%d\n", c.WriteTimeout)
	log.Printf("TLSCipher: \t\t%s\n", c.TLSCipher)
	log.Printf("TLSMinVersion: \t%s\n", c.TLSMinVersion)
	log.Printf("TLSCert: \t\t%s\n", c.TLSCert)
	log.Printf("TLSKey: \t\t%s\n", c.TLSKey)
}
