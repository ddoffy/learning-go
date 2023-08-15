// File: main.go
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// main is the entry point of the program
func main() {
	//r, closer, err := buildGZipReader("my_data.txt.gz")

	start := time.Now()

	r, closer, err := buildGZipReader("sample_data.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	defer closer()

	letters, err := countLetters(r)
	if err != nil {
		panic(err)
	}

	duration := time.Since(start)

	fmt.Printf("Time execution: %v of %v", duration, letters)

}

// countLetters counts the number of letters in a string
// and returns a map of the letters and their counts
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}

	}
}

// buildGZipReader to unzip a file
func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)

	if err != nil {
		return nil, nil, err
	}

	gr, err := gzip.NewReader(r)

	if err != nil {
		return nil, nil, err
	}

	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}
