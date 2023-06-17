package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/dlclark/regexp2"
)

func main() {
	// Open input file
	input, err := os.ReadFile("input-text.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Prepare pattern
	emailPattern := `[\w\.+-]+@[\w\.-]+\.[\w\.-]+`
	uriPattern := `[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`
	ipPattern := `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`

	// Standard library
	stdEmailDuration, stdEmailMatch := measureStdLib(input, emailPattern)
	stdUriDuration, stdUriMatch := measureStdLib(input, uriPattern)
	stdIpDuration, stdIpMatch := measureStdLib(input, ipPattern)

	// Regexp2 library
	rxp2EmailDuration, rxp2EmailMatch := measureRegexp2(input, emailPattern)
	rxp2UriDuration, rxp2UriMatch := measureRegexp2(input, uriPattern)
	rxp2IpDuration, rxp2IpMatch := measureRegexp2(input, ipPattern)

	// Print result
	fmt.Println("STANDARD LIBRARY")
	fmt.Printf("EMAIL: GOT %d MATCHES IN %dms\n", stdEmailMatch, stdEmailDuration)
	fmt.Printf("URI  : GOT %d MATCHES IN %dms\n", stdUriMatch, stdUriDuration)
	fmt.Printf("IP   : GOT %d MATCHES IN %dms\n", stdIpMatch, stdIpDuration)
	fmt.Println()
	fmt.Println("REGEXP2 LIBRARY")
	fmt.Printf("EMAIL: GOT %d MATCHES IN %dms\n", rxp2EmailMatch, rxp2EmailDuration)
	fmt.Printf("URI  : GOT %d MATCHES IN %dms\n", rxp2UriMatch, rxp2UriDuration)
	fmt.Printf("IP   : GOT %d MATCHES IN %dms\n", rxp2IpMatch, rxp2IpDuration)
}

func measureStdLib(data []byte, pattern string) (ms int64, count int) {
	// Compile regex
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	// Run regex
	start := time.Now()
	matches := r.FindAll(data, -1)
	elapsed := time.Since(start)
	return elapsed.Milliseconds(), len(matches)
}

func measureRegexp2(data []byte, pattern string) (ms int64, count int) {
	// Compile regex
	r, err := regexp2.Compile(pattern, regexp2.IgnoreCase)
	if err != nil {
		log.Fatal(err)
	}

	// Run regex
	strData := string(data)
	start := time.Now()

	var matches []string
	m, _ := r.FindStringMatch(strData)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = r.FindNextMatch(m)
	}

	elapsed := time.Since(start)
	return elapsed.Milliseconds(), len(matches)
}
