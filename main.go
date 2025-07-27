package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/SrPatoMan/bakfinder/bakfinderfunctions"
)

func main() {

	target := flag.String("u", "", "URL of your target. Ej: -u https://subdomain.domain.com")
	targetFile := flag.String("l", "", "Subdomain list")
	concurrent := flag.Int("t", 20, "Amount of threads.")
	flag.Parse()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("\n[-] Exiting...")
		os.Exit(0)
	}()

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
			subdomain := strings.TrimSpace(scanner.Text())
			subdomain = strings.TrimSuffix(subdomain, "/")
			permutations := bakfinderfunctions.Permutations(subdomain)

			ch <- struct{}{}
			wg.Add(1)
			go bakfinderfunctions.Fuzzing(subdomain, permutations, ch, &wg)
		}

		wg.Wait()

	}

	if *targetFile == "" && *target != "" {

		ch := make(chan struct{}, *concurrent)
		var wg sync.WaitGroup

		subdomain := strings.TrimSpace(*target)
		subdomain = strings.TrimSuffix(subdomain, "/")

		if !strings.HasPrefix(subdomain, "http://") && !strings.HasPrefix(subdomain, "https://") {
			fmt.Printf("[!] The URL %s is wrong, enter a correct URL. Ex: https://sub1.sub2.target.com\n", subdomain)
			os.Exit(1)
		}

		permutations := bakfinderfunctions.Permutations(subdomain)
		ch <- struct{}{}
		wg.Add(1)
		go bakfinderfunctions.Fuzzing(subdomain, permutations, ch, &wg)

		wg.Wait()

	}
}
