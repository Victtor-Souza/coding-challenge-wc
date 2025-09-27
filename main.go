package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type CountFileResult struct {
	filename   string
	totalLines int64
	totalBytes int
	totalWords int
	totalChars int
}

type CountConfig struct {
	countbytes bool
	countlines bool
	countwords bool
	countchars bool
	countall   bool
}

func countFile(reader io.Reader, filename string) CountFileResult {
	result := CountFileResult{
		filename: filename,
	}

	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, reader)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(bytes.NewReader(buffer.Bytes()))
	for scanner.Scan() {
		line := scanner.Text()
		result.totalLines++
		result.totalWords += len(strings.Fields(line))
	}

	bytes, chars := countBytesAndChar(bytes.NewReader(buffer.Bytes()))

	result.totalBytes = bytes
	result.totalChars = chars

	return result
}

func countBytesAndChar(reader io.Reader) (int, int) {
	bytes, err := io.ReadAll(reader)
	totalBytes := len(bytes)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	chars := 0
	for len(bytes) > 0 {
		_, size := utf8.DecodeRune(bytes)
		chars++
		bytes = bytes[size:]
	}

	return totalBytes, chars
}

func printResult(result CountFileResult, config CountConfig) {
	if config.countall {
		fmt.Printf("%d %d %d %s\n", result.totalLines, result.totalWords, result.totalBytes, result.filename)
		os.Exit(0)
	}

	switch {
	case config.countbytes:
		fmt.Printf("%d %s\n", result.totalBytes, result.filename)
		os.Exit(0)
	case config.countchars:
		fmt.Printf("%d %s\n", result.totalChars, result.filename)
		os.Exit(0)
	case config.countlines:
		fmt.Printf("%d %s\n", result.totalLines, result.filename)
		os.Exit(0)
	case config.countwords:
		fmt.Printf("%d %s\n", result.totalWords, result.filename)
		os.Exit(0)
	}

}

func main() {
	bytes := flag.Bool("c", false, "count the bytes of file")
	lines := flag.Bool("l", false, "count the lines in a file")
	words := flag.Bool("w", false, "count the words in a file")
	chars := flag.Bool("m", false, "count the number of chars in a file")

	flag.Parse()

	config := CountConfig{
		countbytes: *bytes,
		countlines: *lines,
		countwords: *words,
		countchars: *chars,
		countall:   flag.NFlag() == 0,
	}

	if len(flag.Args()) == 0 {
		result := countFile(os.Stdin, "")
		printResult(result, config)

	}

	filename := flag.Args()[0]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	result := countFile(file, filename)
	printResult(result, config)

}
