package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/iam-cedano/fastcrop/internal/images"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("Usage: fastcrop <image-path> <width> <height>")
		os.Exit(1)
	}

	path := args[0]
	width, err := strconv.Atoi(args[1])
	height, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	image := images.Image{Path: path}

	err = image.Crop(width, height)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

}
