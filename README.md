<h1 align="center">Bakfinder</h1>

Bakfinder is a tool that helps you search for exposed backups and configuration files. The tool generates a wordlist based on permutations of your target's hostname or a list of targets.    

## Installation   

You can install the tool with `go install` or download the repo and doing `go build main.go`   

Installing with go install:   

```
go install github.com/SrPatoMan/bakfinder/cmd/bakfinder@latest
```   

## Usage   

Examples of use:   

Fuzzing a wordlist of subdomains with 50 threads      
```
bakfinder -l subdomains.txt -t 50
```   

Fuzzing a single target (with default concurrency)   
```
bakfinder -u "https://subdomain1.subdomain2.domain.com"
```   

## Options   

