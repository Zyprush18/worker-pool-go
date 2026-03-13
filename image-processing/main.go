package main

import (
	"errors"
	"image"
	"os"
	"path/filepath"

	"strings"

	"github.com/gen2brain/avif"
)

func main() {
	dir := "./image"
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	img := make(chan image.Image, len(files))
	name := make(chan string, len(files))
	for _, v := range files {
		f, err := os.Open(filepath.Join(dir, v.Name()))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		fileName := strings.Split(v.Name(), ".")
		name <- fileName[0]
		switch fileName[1] {
		case "jpg", "png", "gif":
			imge, _, err := image.Decode(f)
			if err != nil {
				panic(err)
			}
			img <- imge
		case "avif":
			imge, err := avif.Decode(f)
			if err != nil {
				panic(err)
			}
			img <- imge

		default:
			panic(errors.New("unsupported file format"))
		}
	}

	close(name)
	close(img)

	// convert := NewImageJobConvert(3, img, name, "jpg")
	// convert.WorkerConvert()

	// fmt.Println(<-name)
	resize := NewImageJobResize(3, img,name, 300, 400)
	resize.WorkerResize()
}
