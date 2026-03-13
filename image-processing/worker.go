package main

import (
	"fmt"
	"image"
	"path/filepath"
	"sync"

	"github.com/disintegration/imaging"
)

var (
	convertDir string = "./image-convert"
	resizeDir string = "./image-resize"
	compressDir string = "./image-compress"
)


type ConvertJob struct {
	Name chan string
	Image chan image.Image
	Format string
	Wg sync.WaitGroup
	Mx sync.Mutex
}

type ResizeImage struct {
	Name chan string
	Image chan image.Image
	Width int
	Height int
	Wg sync.WaitGroup
	Mx sync.Mutex
}

type CompressImage struct {
	Name chan string
	Image chan image.Image
	Quality int
	Wg sync.WaitGroup
}

type ImageJob struct {
	Convert ConvertJob
	Resize ResizeImage
	Compress CompressImage
	Worker int
}

func NewImageJobConvert(worker int, img chan image.Image, name chan string, format string) *ImageJob {
	return &ImageJob{
		Worker: worker,
		Convert: ConvertJob{
			Name: name,
			Image: img,
			Format: format,
			Wg: sync.WaitGroup{},
			Mx: sync.Mutex{},
		},
	}
}

func NewImageJobResize(worker int, img chan image.Image, name chan string,width, height int) *ImageJob {
	return &ImageJob{
		Worker: worker,
		Resize: ResizeImage{
			Name: name,
			Image:  img,
			Width:  width,
			Height: height,
			Wg:     sync.WaitGroup{},
			Mx: sync.Mutex{},
		},
	}
}

func NewImageJobCompress(worker int, img chan image.Image, name chan string,quality int) *ImageJob {
	return &ImageJob{
		Worker: worker,
		Compress: CompressImage{
			Name: name,
			Image:   img,
			Quality: quality,
			Wg:      sync.WaitGroup{},
		},
	}
}

func (j *ImageJob) WorkerConvert() {
	for i := 1; i <= j.Worker; i++ {
		j.Convert.Wg.Add(1)
		go j.ConvertProcess(i)
	}

	j.Convert.Wg.Wait()
}


func (j *ImageJob) WorkerResize() {
	for i := 1; i <= j.Worker; i++ {
		j.Resize.Wg.Add(1)
		go j.ResizeProcess(i)
	}
	j.Resize.Wg.Wait()
}

func (j *ImageJob) WorkerCompress() {
	for i := 1; i <= j.Worker; i++ {
		j.Compress.Wg.Add(1)
		go j.CompressProcess(i)
	}

	j.Compress.Wg.Wait()
}



func (j *ImageJob) ConvertProcess(i int) {
	defer j.Convert.Wg.Done()
	for v := range j.Convert.Image {
		fmt.Printf("Worker %d sedang kerja\n", i)
		if err:= imaging.Save(v, filepath.Join(convertDir, <-j.Convert.Name+"."+j.Convert.Format));err != nil {
			panic(err)
		}
	}
}


func (j *ImageJob) ResizeProcess(i int) {
	defer j.Resize.Wg.Done()
	for v := range j.Resize.Image {
		name := <-j.Resize.Name+".jpg"
		fmt.Printf("worker %d riseze image %s\n", i, name)
		img:= imaging.Resize(v, j.Resize.Width, j.Resize.Height, imaging.NearestNeighbor)
		if err:= imaging.Save(img, filepath.Join(resizeDir, name));err != nil {
			panic(err)
		}
	}
}


func (j *ImageJob) CompressProcess(i int) {
	
}