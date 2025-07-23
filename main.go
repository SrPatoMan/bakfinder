package main

import (
	"bakfinder/bakfinderfunctions"
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {

	target := flag.String("u", "", "URL of your target. Ej: -u https://subdomain.domain.com")
	targetFile := flag.String("l", "", "Target list")
	concurrent := flag.Int("t", 20, "Amount of threads")
	flag.Parse()

	if *target == "" && *targetFile == "" {
		fmt.Println("[!] Provide a subdomain input file (-l) or a single target (-u)")
		os.Exit(1)
	}

	if *target != "" && *targetFile != "" {
		fmt.Println("[!] Both parameters are not allowed, provide either a subdomains file (-l) or a single subdomain (-u)")
		os.Exit(1)
	}

	if *target == "" && *targetFile != "" {
		subdomainsFile, err := os.Open(*targetFile)
		if err != nil {
			fmt.Println("[!] Error reading the subdomains file...")
			os.Exit(1)
		}
		defer subdomainsFile.Close()

		scanner := bufio.NewScanner(subdomainsFile)
		ch := make(chan struct{}, *concurrent)
		var wg sync.WaitGroup

		for scanner.Scan() {
			subdomain := scanner.Text()
			permutations := bakfinderfunctions.Permutations(subdomain)

			ch <- struct{}{}
			wg.Add(1)
			go bakfinderfunctions.Fuzzing(subdomain, permutations, ch, &wg)
		}

		wg.Wait()

	}
}
