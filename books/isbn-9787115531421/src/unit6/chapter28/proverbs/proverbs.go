package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	err := writeProverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readProverbs(filename string) []string {
	var content []string
	data, err := os.ReadFile(filename)
	if err != nil {
		return content
	}

	content = append(content, strings.Split(string(data), "\n")...)
	
	return content

}

func writeProverbs(name string) error {
	f, err := os.Create(name)

	if err != nil {
		return err
	}

	defer f.Close()

	sw := safeWriter{writer: f}
	completeProverbs := readProverbs("proverbs-complete.txt")
	for _, p := range completeProverbs {
		sw.writeln(p)
	}
	return sw.err
}