package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	target := flag.String("u", "", "URL of your target. Ej: -u https://subdomain.domain.com")
	target_file := flag.String("l", "", "Target list")
	flag.Parse()

	if *target == "" && *target_file == "" {
		fmt.Println("[!] Provide a subdomain input file (-l) or a single target (-u)")
		os.Exit(1)
	}

}
