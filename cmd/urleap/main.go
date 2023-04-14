package main

import (
	"fmt"
	"os"
)

func goMain(args []string) int {
	fmt.Println("Hello World")
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
