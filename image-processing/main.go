package main

import (
	"fmt"
	"log"
	"os"

	"path/filepath"

	"github.com/h2non/bimg"
	"strings"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Please added argument (convert/resize/compress)")
	}

	dir := "./image"

	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	img := make(chan []byte, len(files))
	name := make(chan string, len(files))
	for _, v := range files {
		fileName := strings.Split(v.Name(), ".")
		name <- fileName[0]
		buffer, err := bimg.Read(filepath.Join(dir, v.Name()))
		if err != nil {
			panic(err)
		}
		img <- buffer
	}

	close(img)
	close(name)

	JobImage := NewJobImageProcessing(3, name, img, bimg.JPEG)
	switch args[1] {
	case "convert":
		JobImage.WorkerConvert()
	case "resize":
		JobImage.WorkerResize(300, 400)
	case "compress":
		JobImage.WorkerCompress(80, 2048)
	default:
		fmt.Println("Invalid argument. Use: convert/resize/compress")
	}

}
