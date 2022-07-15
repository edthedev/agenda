package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func getAgenda(fileName string) []string {

	agendaRegex, _ := regexp.Compile(`(^[\+-]\W[\d:.]+\W[ap]m.+$)`) // How we find agenda items.
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Unable to read file: %s - %s", fileName, err)
	}
	scanner := bufio.NewScanner(file)
	var results = []string{}
	var result = []string{}
	for scanner.Scan() {
		result = agendaRegex.FindAllString(scanner.Text(), -1)
		results = append(results, result...)
	}

	return results
}

func main() {
	code := 0
	defer func() {
		os.Exit(code)
	}()
	agendaFile := flag.String("path", "", "Search this file agenda lines.")
	// maxFlag := flag.Int("max", 5, "Maximum number of lines to list.")
	flag.Parse()

	var found = []string{}
	found = getAgenda(*agendaFile)
	// TODO: Add a flag the removes line breaks, to allow files to be passed to Vim

	var item string
	for _, item = range found {
		fmt.Println(item)
		fmt.Println("")
	}
}
