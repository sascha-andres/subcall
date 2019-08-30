package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var executableName string

func init() {
	fileWithPath, _ := os.Executable()
	executableName = filepath.Base(fileWithPath)
}

func help() {
	fmt.Println(fmt.Sprintf("Usage %s <command>", executableName))
}

func main() {
	if len(os.Args) < 2 || strings.ToLower(os.Args[1]) == "--help" {
		help()
		os.Exit(0)
	}

	var result int
	path, err := exec.LookPath(fmt.Sprintf("%s-%s", executableName, os.Args[1]))
	if err != nil {
		fmt.Printf("error: %s\n", err)
		result = 1
	} else {
		cmd := exec.Command(path)
		cmd.Args = os.Args[1:]
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		if err := cmd.Start(); err != nil {
			fmt.Println(fmt.Sprintf("error: unable to start [%s]: %s", path, err.Error()))
			result = 1
		}
		if err := cmd.Wait(); err != nil {
			result = 1
		}
	}
	os.Exit(result)
}
