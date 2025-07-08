package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("ehllo")
	cmds := os.Args

	if len(cmds) < 3 {
		fmt.Println("Not enough parmeters")
		return
	}
	// fmt.Println(len(cmds))
	matcher := cmds[1]
	fileName := cmds[len(cmds)-1]
	if !strings.Contains(fileName, ".txt") {
		fmt.Println("Give correct filename")
		return
	}
	// fmt.Println(matcher)
	// fmt.Println(fileName)
	if printMatchingLines(matcher, fileName) != nil {
		fmt.Println("Error in printing")
		return
	}
}

func printMatchingLines(matcher, fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// file.Close()
	fileContent := string(file)
	lines := strings.Split(fileContent, "\n")
	// fmt.Println(len(lines))
	for _, r := range lines[:len(lines)-1] {
		if strings.Contains(r, matcher) {
			fmt.Println(r)
		}
	}
	return nil
}
