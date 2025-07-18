package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
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

	var paths []string

	paths, err = getDirectoryName(cmds[1:], matcher)

	para, err := getCommands(cmds[1:])
	fmt.Println("Directory", paths)
	printDirOrFile(paths)
	fmt.Println("Parameter", para)
	if len(para) == 0 {
		traverseFiles(matcher, paths)
		// printMatchingLines(matcher,)
	}
}

func traverseFiles(matcher string, fileNames []string) error {
	for _, r := range fileNames {
		isDir, err := isDir(r)
		if err != nil || isDir {
			return errors.New("File error")
		}
		printMatchingLines(matcher, r)
	}
	return nil
}

func traversePath(paths []string) error {

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
			fmt.Println(fileName, " : ", r)
		}
	}
	return nil
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
		if !strings.HasPrefix(r, "-r") && !(matcher == r) {
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

func isDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func printDirOrFile(path []string) {
	for _, r := range path {
		isDir, err := isDir(r)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(r, isDir)
	}
}
