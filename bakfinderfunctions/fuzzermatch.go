package bakfinderfunctions

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func Fuzzing(subdomain string, payloads []string, ch chan struct{}, wg *sync.WaitGroup) {
	defer func() { <-ch }()
	defer wg.Done()

	controlUrl := fmt.Sprintf("%s/", subdomain)
	controlRequest, controlErr := http.Get(controlUrl)
	if controlErr != nil {
		return
	}

	controlBody, controlErrBody := io.ReadAll(controlRequest.Body)
	if controlErrBody != nil {
		controlBody = []byte{}
	}
	controlLength := len(controlBody)

	for _, payload := range payloads {

		myurl := fmt.Sprintf("%s/%s", subdomain, payload)
		resp, err := http.Get(myurl)
		if err != nil {
			return
		}

		var earlyReturn bool

		func() {
			location := resp.Header.Get("Location")
			locationParsing, parsingErr := url.Parse(location)
			if parsingErr != nil {
				return
			}
			locationHostname := locationParsing.Hostname()

			hostname := strings.Split(subdomain, ".")
			hostname = hostname[len(hostname)-2:]
			hostnameSubdomain := strings.Join(hostname, ".")

			if location != "" && !strings.HasSuffix(locationHostname, hostnameSubdomain) {
				earlyReturn = true
			}
		}()

		if earlyReturn {
			return
		}

		if resp.StatusCode != 200 {
			return
		}

		body, bodyErr := io.ReadAll(resp.Body)
		if bodyErr != nil {
			return
		}
		defer resp.Body.Close()

		length := len(body)

		if resp.StatusCode == 200 && length != controlLength {
			fmt.Printf("\033[32m[+]\033[0m %s found\n", myurl)
		}

	}

}
