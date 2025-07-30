package bakfinderfunctions

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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

	falsePositivePatterns := []string{
		"window.location",
		"location =",
		"location=",
		"location.href",
		"window.location.href",
		"location.assign",
		"location.replace",
		"document.forms[0].submit()",
		"blazor.webassembly.js",
		"_framework/blazor",
		"blazor-environment",
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
			continue
		}

		body, bodyErr := io.ReadAll(resp.Body)
		if bodyErr != nil {
			continue
		}
		resp.Body.Close()

		bodyStr := strings.ToLower(string(body))

		for _, falsePositivePattern := range falsePositivePatterns {
			if strings.Contains(bodyStr, falsePositivePattern) {
				continue
			}
		}

		length := len(body)

		if resp.StatusCode == 200 && length != controlLength {
			fmt.Printf("\033[32m[+]\033[0m %s \033[32mfound\033[0m\n", myurl)
		}

	}

}
