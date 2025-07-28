package bakfinderfunctions

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func Fuzzing(subdomain string, payloads []string, ch chan struct{}, wg *sync.WaitGroup) {
	defer func() { <-ch }()
	defer wg.Done()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	controlUrl := fmt.Sprintf("%s/", subdomain)
	controlRequest, controlErr := client.Get(controlUrl)
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

		resp, err := client.Get(myurl)
		if err != nil {
			return
		}

		body, bodyErr := io.ReadAll(resp.Body)
		if bodyErr != nil {
			return
		}
		defer resp.Body.Close()

		length := len(body)

		fmt.Printf("[DEBUG] url: %s | status: %d | length: %d | control: %d\n", myurl, resp.StatusCode, length, controlLength)

		if resp.StatusCode == 200 && length != controlLength {
			fmt.Printf("\033[32m[+]\033[0m %s found\n", myurl)
		}

	}

}
