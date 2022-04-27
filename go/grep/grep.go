package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	allMatches := []string{}
	flagMap := map[string]bool{}
	for _, flag := range flags {
		flagMap[flag] = true
	}
	for _, filename := range files {
		contents, err := ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		isMatchSlice := make([]bool, len(contents))
		for index, line := range contents {
			isMatchSlice[index] = IsMatch(pattern, line, flagMap)
		}
		fileMatches := []string{}
		for index, isMatch := range isMatchSlice {
			if isMatch {
				outputLine := ""
				if len(files) > 1 {
					outputLine += fmt.Sprintf("%s:", filename)
				}
				if flagMap["-n"] {
					outputLine += fmt.Sprintf("%d:", index+1)
				}
				outputLine += contents[index]
				fileMatches = append(fileMatches, outputLine)
			}
		}
		if len(fileMatches) > 0 && flagMap["-l"] {
			allMatches = append(allMatches, filename)
			continue
		}
		allMatches = append(allMatches, fileMatches...)
	}
	return allMatches
}

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	contents := []string{}
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	return contents, nil
}

func IsMatch(pattern, line string, flags map[string]bool) bool {
	var result bool
	if flags["-i"] {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}
	if flags["-x"] {
		result = line == pattern
	} else {
		result = strings.Contains(line, pattern)
	}
	if flags["-v"] {
		result = !result
	}
	return result
}
