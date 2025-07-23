package bakfinderfunctions

import (
	"fmt"
	"net/url"
	"strings"
)

func Permutations(subdomain string) []string {

	target, err := url.Parse(subdomain)

	if err != nil {
		fmt.Printf("[!] Error: %v", subdomain)
		return []string{}
	}

	extensions := []string{"bak", "zip", "7z", "rar", "old", "backup", "orig", "original", "src", "dev", "inc", "copy", "tmp", "swp", "tar", "gz", "tar.gz", "sql", "db", "bd", "ddbb", "bbdd", "csv", "mdb", "accdb", "dbs", "sqlite", "frm", "ibd", "mwb", "myd", "mrg", "tmd", "xml", "json", "yaml", "yml", "env", "ini", "cfg", "log", "md", "txt", "xls", "xlsx", "doc", "docs", "docx", "pdf", "conf", "config", "php.old", "py", "rb", "jar", "jar.old", "java", "old.jsp", "jsp.old"}

	permutations := []string{}

	hostname := target.Hostname()
	hostname_parts := strings.Split(hostname, ".")

	if len(hostname_parts) <= 2 {
		return []string{}
	}

	subdomains := hostname_parts[:len(hostname_parts)-2]
	domain := hostname_parts[len(hostname_parts)-2:]
	domain_without_tld := hostname_parts[len(hostname_parts)-2 : len(hostname_parts)-1]
	all_without_tld := hostname_parts[:len(hostname_parts)-1]

	for _, sub := range subdomains {
		for _, ext := range extensions {
			permutation := fmt.Sprintf("%s.%s", sub, ext)
			permutations = append(permutations, permutation)
		}
	}

	for _, domain := range domain {
		for _, ext := range extensions {
			permutation := fmt.Sprintf("%s.%s", domain, ext)
			permutations = append(permutations, permutation)
		}
	}

	for _, domain := range domain_without_tld {
		for _, ext := range extensions {
			permutation := fmt.Sprintf("%s.%s", domain, ext)
			permutations = append(permutations, permutation)
		}
	}

	separators := []string{".", "_", "-"}

	for _, separator := range separators {
		subdomainsStr := strings.Join(subdomains, separator)
		for _, ext := range extensions {
			permutation := fmt.Sprintf("%s.%s", subdomainsStr, ext)
			permutations = append(permutations, permutation)
		}
	}

	for _, separator := range separators {
		subdomainsStr := strings.Join(hostname_parts, separator)
		for _, ext := range extensions {
			permutation := fmt.Sprintf("%s.%s", subdomainsStr, ext)
			permutations = append(permutations, permutation)
		}
	}

	for _, separator := range separators {
		subdomainsStr := strings.Join(all_without_tld, separator)
		for _, ext := range extensions {
			permutation := fmt.Sprintf("%s.%s", subdomainsStr, ext)
			permutations = append(permutations, permutation)
		}
	}

	return permutations
}
