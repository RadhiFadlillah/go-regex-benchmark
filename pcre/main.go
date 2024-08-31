package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	pcre "github.com/GRbit/go-pcre"
)

func measure(data []byte, pattern string) {
	r, err := pcre.CompileJIT(pattern, 0, pcre.STUDY_JIT_COMPILE)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	count := len(r.FindAllIndex(data, 0))
	elapsed := time.Since(start)

	fmt.Printf("%f - %v\n", float64(elapsed)/float64(time.Millisecond), count)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: benchmark <filename>")
		os.Exit(1)
	}

	filerc, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	data := buf.Bytes()

	// Email
	measure(data, `[\w\.+-]+@[\w\.-]+\.[\w\.-]+`)

	// URI
	measure(data, `[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`)

	// IP
	measure(data, `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`)

	// Named groups
	//measure(data, `(?P<email>[\w\.+-]+@[\w\.-]+\.[\w\.-]+)|(?P<uri>[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?)|(?P<ip>(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]))`)
}
