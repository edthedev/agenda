package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
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

func getAgendaFile(agendaDate time.Time, format string) string {
	year, month, day := agendaDate.Date()
	var result = strings.Replace(format, "{YYYY}", fmt.Sprintf("%d", year), 1)
	result = strings.Replace(result, "{MM}", fmt.Sprintf("%02d", month), 1)
	result = strings.Replace(result, "{DD}", fmt.Sprintf("%02d", day), 1)
	return result
}

func main() {
	code := 0
	defer func() {
		os.Exit(code)
	}()
	agendaFile := flag.String("path", "", "Search this file agenda lines.")
	flag.Parse()

	fileName := getAgendaFile(time.Now(), *agendaFile)

	var found = []string{}
	found = getAgenda(fileName)

	var item string
	for _, item = range found {
		fmt.Println(item)
		fmt.Println("")
	}
}
