package bakfinderfunctions

import (
	"net/http"
	"time"
)

func HttpClient() *http.Client {

	transport := &http.Transport{
		MaxConnsPerHost:       20,
		MaxIdleConns:          500,
		DisableCompression:    true,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   8 * time.Second,
		ResponseHeaderTimeout: 8 * time.Second,
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: transport,
	}

	return client
}
