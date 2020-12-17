package main

import (
	"fmt"
	"os/exec"
	"os"
	"time"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func dirExists(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

// Check for dir path
// If path does not exist, go and create new one
func getOrCreateDirPath(dir string) string {
	home, err := os.UserHomeDir()
	check(err)

	var path string
	if dir == "" {
		path = fmt.Sprintf("%s/notes", home)
	} else {
		path = fmt.Sprintf("%s/notes/%s", home, dir)
	}

	if dirExists(path) {
		return path
	}

	err = os.MkdirAll(path, 0777)
	check(err)

	return path
}

// Get file base on directory
// If directory path doesn't exist, create new one.
// If file doesn't exist, create new one
func getOrCreateFilePath(dir string, file string) string {
	path := fmt.Sprintf("%s/%s", getOrCreateDirPath(dir), file)
	if fileExists(path) {
		return path
	}

	_, err := os.Create(path)
	check(err)

	return path
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDefaultFileName() string {
	currentDate := time.Now().Local()
	fileName := fmt.Sprintf("%d%d%d.md", currentDate.Day(), currentDate.Month(), currentDate.Year())

	return fileName
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var fileName string
	if len(os.Args) <= 1 {
		fileName = getDefaultFileName()
	} else {
		fileName = os.Args[1]
	}
	path := getOrCreateFilePath("", fileName)

	cmd := exec.Command("nvim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		panic(err.Error())
	}
}
