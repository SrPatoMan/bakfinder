package bakfinderfunctions

import (
	"fmt"
	"io"
	"net/http"
)

func fuzzing(subdomain string, payloads []string) {

	controlUrl := fmt.Sprintf("%s/", subdomain)
	controlRequest, _ := http.Get(controlUrl)

	controlBody, _ := io.ReadAll(controlRequest.Body)
	controlLength := len(controlBody)

	for _, payload := range payloads {

		url := fmt.Sprintf("%s/%s", subdomain, payload)
		resp, _ := http.Get(url)

		body, _ := io.ReadAll(resp.Body)
		length := len(body)

		if resp.StatusCode == 200 && length != controlLength {
			fmt.Printf("[+]%s encontrado", url)
		}

	}

}
