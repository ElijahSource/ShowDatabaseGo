package main

import (
	"fmt"
	"os"
)

func main() {
	dirname := "E:/Tv Shows/Comedy Shows" // replace with your directory

	var Shows [0]string
	files, err := os.ReadDir(dirname)
	type mytype struct {
		a, b int
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		Shows = append(Shows, file.Name())
	}
}
