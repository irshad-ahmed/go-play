package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		fileName := os.Args[1]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		} else {
			io.Copy(os.Stdout, file)
		}
	} else {
		fmt.Println("File Name expected as part of first paramenter  ")
	}
}
