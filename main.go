package main

import (
	"errors"
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
	matcher, err := getMatchingString(cmds[1:])
	fmt.Println("Matcher:", matcher)
	if err != nil {
		fmt.Println("Please give a matcing string")
		return
	}
	hasFile := true
	hasDir := true
	hasPara := true
	fileName, err := getFileName(cmds[1:])
	if err != nil {
		hasFile = false
	}
	var dirName []string
	if !hasFile {
		dirName, err = getDirectoryName(cmds[1:], matcher)
		if err != nil {
			hasDir = false
		}
	}
	if !hasDir && !hasFile {
		fmt.Println("Does not have file or director name")
		return
	}

	para, err := getCommands(cmds[1:])
	if err != nil {
		hasPara = false
	}
	fmt.Println("Directory", dirName)
	fmt.Println("Filename", fileName)
	// fmt.Println("Parameters", para, hasPara)

	if hasFile && !hasPara {
		if printMatchingLines(matcher, fileName) != nil {
			fmt.Println("Error in printing")
			return
		}
	} else if hasDir && hasPara {
		if traverseAndPrint(dirName, para) != nil {
			fmt.Println("Error in printing")
			return
		}
	} else {
		fmt.Println("Incorrect parameters")
		return
	}

	// fmt.Println(matcher)
	// fmt.Println(fileName)
	// if printMatchingLines(matcher, fileName) != nil {
	// 	fmt.Println("Error in printing")
	// 	return
	// }
}

func traverseAndPrint(dir []string, para []string) error {
	return nil
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

func getFileName(cmds []string) (string, error) {
	for _, r := range cmds {
		if strings.Contains(r, ".txt") {
			return r, nil
		}
	}
	return "", errors.New("Could not find text file")
}

func getCommands(cmds []string) ([]string, error) {
	resList := make([]string, 0)
	for _, r := range cmds {
		if strings.HasPrefix(r, "-") {
			resList = append(resList, r)
		}
	}
	if len(resList) == 0 {
		return nil, errors.New("No command present")
	}
	return resList, nil
}

func getMatchingString(cmds []string) (string, error) {
	for _, r := range cmds {
		if !strings.Contains(r, "-") && !strings.Contains(r, ".txt") {
			return r, nil
		}
	}

	return "", errors.New("No matching string")
}

func getDirectoryName(cmds []string, matcher string) ([]string, error) {
	// fmt.Println(cmds)
	dirs := make([]string, 0)
	for _, r := range cmds {
		// fmt.Println("outside condition:", r)
		if !strings.HasPrefix(r, "-r") && !strings.Contains(r, ".txt") && !(matcher == r) {
			// fmt.Println("inside condition", r)
			// return r, nil
			dirs = append(dirs, r)
		}
		// fmt.Println(dirs)
	}
	if len(dirs) == 0 {
		return nil, errors.New("No matching string")
	}

	return dirs, nil
}
