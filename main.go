package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("Usage: fastcrop <image-path> <width> <height>")
		os.Exit(1)
	}

	_ = args[0]
	if _, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	if _, err := strconv.Atoi(args[2]); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
