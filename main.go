package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func countWords(content []byte) int {
	return len(strings.Fields(string(content)))
}

func countLines(content []byte) int {
	return len(strings.Split(strings.Trim(string(content), "\n\\s"), "\n"))
}

func main() {
	bytes := flag.Bool("c", false, "Count bytes")
	lines := flag.Bool("l", false, "Count lines")
	words := flag.Bool("w", false, "Count words")
	local := flag.Bool("m", false, "Count locals")
	flag.Parse()
	fileName := flag.Arg(0)
	var content []byte
	var err error
	switch fileName {
	case "":
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			content = append(content, scanner.Bytes()...)
		}
	default:
		content, err = os.ReadFile(fileName)
	}
	if err != nil {
		log.Fatal(err)
	}
	if *bytes {
		fmt.Printf("%d %s", len(content), fileName)
	}
	if *lines {
		fmt.Printf("%d %s", countLines(content), fileName)
	}
	if *words {
		fmt.Printf("%d %s", countWords(content), fileName)
	}
	if *local {
		fmt.Printf("%d %s", utf8.RuneCount(content), fileName)
	}
	if flag.NFlag() == 0 {
		fmt.Printf("  %d  %d %d %s", countLines(content), countWords(content), len(content), fileName)
	}
}
