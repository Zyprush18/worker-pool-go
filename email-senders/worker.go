package main

import (
	"fmt"
	"sync"
)

type EmailSendJob struct {
	worker int
	email chan EmailsTask
	wg 	sync.WaitGroup
	mx sync.Mutex
}

var workerKerja int

func NewWorker(workerID int, emailChan chan EmailsTask) *EmailSendJob {
	return &EmailSendJob{
		worker: workerID,
		email: emailChan,
		wg: sync.WaitGroup{},
		mx: sync.Mutex{},
	}
}

func (e *EmailSendJob) Worker()  {
	for i := 1; i <= e.worker; i++ {
		e.wg.Add(1)
		go e.Process(i)
	}

	e.wg.Wait()
}

func (e *EmailSendJob) Process(i int) {
	defer e.wg.Done()
	for email := range e.email {
		e.mx.Lock()
		fmt.Printf("Worker %d processing email: %s and sending to %s with subject %s and message: %s\n", i, email.From.Email, email.To.Email, email.Subject, email.Body)
		workerKerja++
		e.mx.Unlock()
	}
}