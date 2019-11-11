package pingdom

import (
	"log"
	"net/http"
	"time"
)

var host = "<host-address>"

// PingdomAPIURL ...
const PingdomAPIURL = "https://api.pingdom.com/2.1"

// Username ...
var Username = "<username>"

// Password ...
var Password = "<password>"

// AccountEmail ...
var AccountEmail = "<email-address>"

// AppKey ...
var AppKey = "<app-key>"

// PingdomAppKey ...
var PingdomAppKey = "<app-key>"

// Transport ...
type Transport interface {
	Do(*http.Request) (*http.Response, error)
}

// Config ...
type Config struct {
	AppKey       string
	Username     string
	Password     string
	AccountEmail string
	URL          string
	Logger       *log.Logger
	Transport    Transport
}

// ClientConfig ...
type ClientConfig struct {
	Kubernetes struct {
		InCluster bool
		Address   string
	}
	Pingdom struct {
		Username     string
		Password     string
		AccountEmail string
		Timeout      time.Time
	}
}
