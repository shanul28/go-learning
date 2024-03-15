package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	cmd := os.Args
	file, err := os.Open(cmd[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
