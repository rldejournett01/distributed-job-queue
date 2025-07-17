package main

import (
	"log"
	"github.com/rldejournett01/distributed-job-queue/pkg/dispatcher"
)

func main() {
	log.Println("Starting job controller...")
	dispatcher.StartController()
}