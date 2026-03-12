package main

import (
	// "fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Job struct {
	worker      int
	url         chan string
	resultFetch chan ResultFetch
	resultParse chan ResultParse
	Fetchwg     sync.WaitGroup
	Parsewg     sync.WaitGroup
}

type ResultParse struct {
	url string
	title string
}
type ResultFetch struct {
	url string
	html string
}

var (
	berapaFetch int
	berapaParse int
	mx     sync.Mutex
)

func NewJob(worker int, url chan string) *Job {
	return &Job{
		worker:      worker,
		url:         url,
		resultFetch: make(chan ResultFetch),
		resultParse: make(chan ResultParse),
		Fetchwg:     sync.WaitGroup{},
		Parsewg:     sync.WaitGroup{},
	}
}

func (j *Job) WorkerFetchPage() {
	for i := 1; i <= j.worker; i++ {
		j.Fetchwg.Add(1)
		go j.FetchPage(i)
	}
	go func() {
		j.Fetchwg.Wait()
		close(j.resultFetch)

	}()
}

func (j *Job) FetchPage(worker int) {
	defer j.Fetchwg.Done()
	for v := range j.url {
		mx.Lock()
		berapaFetch++
		mx.Unlock()
		res, err := http.Get(v)
		if err != nil {
			log.Fatal(err.Error())
		}

		content, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		res.Body.Close()
		j.resultFetch <- ResultFetch{
			url: v,
			html: string(content),
		}
	}
}

func (j *Job) ParseResultsPage() {
	for i := 0; i < j.worker; i++ {
		j.Parsewg.Add(1)
		go j.ParsePage(i)
	}
	go func() {
		j.Parsewg.Wait()
		close(j.resultParse)
	}()
}

func (j *Job) ParsePage(i int) {
	defer j.Parsewg.Done()
	var title string
	for v := range j.resultFetch {
		mx.Lock()
		berapaParse++
		mx.Unlock()

		doc, err := html.Parse(strings.NewReader(v.html))
		if err != nil {
			log.Fatal(err.Error())
		}

		for n := range doc.Descendants() {
			if n.Type == html.ElementNode && n.Data == "title" {
				if n.FirstChild != nil {
					title = n.FirstChild.Data
				}
			}
		}
		j.resultParse <- ResultParse{
			url:   v.url,
			title: title,
		}
	}

}
