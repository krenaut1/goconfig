package goconfig

import (
	"os"
	"testing"
)

// Config this structure defines the application properties
type Config struct {
	Profile          string // should only be defined in the config.global.json
	ServerAddr       string
	ServerPort       int
	ClientID         string
	ClientSecret     string
	TokenEndPoint    string
	CertEndPoint     string
	UserInfoEndPoint string
	Path             string
}

var config Config

func TestGoconfig(t *testing.T) {
	config := Config{}
	config.Path = os.Getenv("PATH")
	println(config.Path)
	err := GoConfig(&config)
	if err != nil {
		t.Errorf("Err thrown: %v", err)
	}
	println(config.Profile)
	println(config.ServerAddr)
	println(config.ServerPort)
	println(config.ClientID)
	println(config.ClientSecret)
	println(config.TokenEndPoint)
	println(config.CertEndPoint)
	println(config.UserInfoEndPoint)
}
