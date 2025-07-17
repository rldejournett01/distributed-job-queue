package main

import (
	"log"
	"distributed-job-queue/pkg/dispatcher"
)

func main() {
	log.Println("Starting job controller...")
	dispatcher.StartController()
}