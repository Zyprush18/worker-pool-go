package main

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/h2non/bimg"
)

var (
	convertDir string = "./image-convert"
	resizeDir string = "./image-resize"
	compressDir string = "./image-compress"
)


type ConvertJob struct {
	Wg sync.WaitGroup
}

type ResizeImage struct {
	Width int
	Height int
	Wg sync.WaitGroup
	Mx sync.Mutex
}

type CompressImage struct {
	Quality int
	Size int
	Wg sync.WaitGroup
}

type ImageJob struct {
	Name chan string
	Image chan []byte
	Format bimg.ImageType
	Convert ConvertJob
	Resize ResizeImage
	Compress CompressImage
	Worker int
}

func (j *ImageJob) SelectExtension() string {
	switch j.Format {
	case bimg.JPEG:
		return "jpg"
	case bimg.PNG:
		return "png"
	case bimg.WEBP:
		return "webp"
	default:
		return "unknown"
	}
}

func NewJobImageProcessing(worker int, name chan string, image chan []byte, format bimg.ImageType) *ImageJob {
	return &ImageJob{
		Name:   name,
		Image:  image,
		Format: format,
		Worker: worker,
		Convert: ConvertJob{
			Wg: sync.WaitGroup{},

		},
		Resize: ResizeImage{
			Wg: sync.WaitGroup{},
			Mx: sync.Mutex{},
		},
		Compress: CompressImage{
			Wg: sync.WaitGroup{},
		},
	}
}


func (j *ImageJob) SaveToDir(dir string, name string, img []byte) {
	if err := bimg.Write(filepath.Join(dir, name), img); err != nil {
		panic(err)
	}
}

func (j *ImageJob) WorkerConvert() {
	for i := 1; i <= j.Worker; i++ {
		j.Convert.Wg.Add(1)
		go j.ConvertProcess(i)
	}

	j.Convert.Wg.Wait()
}


func (j *ImageJob) WorkerResize(width, height int) {
	j.Resize.Width = width
	j.Resize.Height = height

	for i := 1; i <= j.Worker; i++ {
		j.Resize.Wg.Add(1)
		go j.ResizeProcess(i)
	}
	j.Resize.Wg.Wait()
}

func (j *ImageJob) WorkerCompress(quality int, size int) {
	j.Compress.Quality = quality

	for i := 1; i <= j.Worker; i++ {
		j.Compress.Wg.Add(1)
		go j.CompressProcess(i)
	}

	j.Compress.Wg.Wait()
}



func (j *ImageJob) ConvertProcess(i int) {
	defer j.Convert.Wg.Done()
	for v := range j.Image {
		fmt.Printf("Worker %d sedang kerja\n", i)
		newImg, err := bimg.NewImage(v).Convert(j.Format)
		if err != nil {
			panic(err)
		}

		name := <-j.Name + "." + j.SelectExtension()

		j.SaveToDir(convertDir, name, newImg)
	}
}


func (j *ImageJob) ResizeProcess(i int) {
	defer j.Resize.Wg.Done()
	for v := range j.Image {
		name := <-j.Name + "." + j.SelectExtension()

		fmt.Printf("worker %d risize image %s\n", i, name)
		newImg, err := bimg.NewImage(v).Resize(j.Resize.Width, j.Resize.Height)
		if err != nil {
			panic(err)
		}

		j.SaveToDir(resizeDir, name, newImg)
	}
}


func (j *ImageJob) CompressProcess(i int) {
	defer j.Compress.Wg.Done()
	for v := range j.Image {
		name := <-j.Name + "." + j.SelectExtension()
		fmt.Printf("worker %d compressing image %s\n", i, name)
		newImg, err := bimg.NewImage(v).Process(bimg.Options{Quality: j.Compress.Quality})
		if err != nil {
			panic(err)
		}
		if len(newImg) > j.Compress.Size {
			j.Compress.Quality -= 5
		}

		j.SaveToDir(compressDir, name, newImg)
	}
}