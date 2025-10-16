package utilis

import (
	"net"
	"net/http"
	"time"
)

var HttpClient = &http.Client{
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second, //wait time for TCP Timeout
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second, //wait time for  TSL Timeout
		ResponseHeaderTimeout: 5 * time.Second, // wait time timeout for response header
		IdleConnTimeout: 90 * time.Second, //wait time timeout after 90seconds idle connection
	},
	Timeout: 0,
}