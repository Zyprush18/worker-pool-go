package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type fromEmail struct {
	Name string `json:"name"`
	Email string `json:"email"`
}
type toEmail struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type EmailsTask struct {
	Id int `json:"id"`
	From fromEmail `json:"from"`
	To toEmail `json:"to"`
	Subject string `json:"subject"`
	Body string `json:"body"`
}

func main() {
	file, err := os.ReadFile("emails.json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	var emails []EmailsTask
	if err := json.Unmarshal(file, &emails); err != nil {
		log.Println(err.Error())
	}

	chanEmail := make(chan EmailsTask, len(emails))
	for _, email := range emails {
		chanEmail <- email
	}
	
	close(chanEmail)
	worker := NewWorker(5, chanEmail)
	worker.Worker()
	

	fmt.Printf("Total worker yang kerja: %d\n", workerKerja)
	fmt.Println("All task email succes send!")


}