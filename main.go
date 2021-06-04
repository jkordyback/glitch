package main

import (
	"fmt"
	"os"
)

type exitCode int

const (
	exitOK exitCode = 0
	//exitError  exitCode = 1
	//exitCancel exitCode = 2
)

func main() {
	code := run()
	os.Exit(int(code))
}

func run() exitCode {
	fmt.Println("Hello Glitch")

	return exitOK
}
