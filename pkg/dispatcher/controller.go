package dispatcher

import (
	"github.com/rldejournett01/distributed-job-queue/cmd/worker"
	"fmt"
)

func StartController() {
	jobChan := make(chan worker.Job,10)

	//Start 3 simulated worker threads
	for i := 1; i <= 3; i++ {
		go worker.Worker(i, jobChan)	
	}

	//Submit Jobs
	for j := 1; j <= 5; j++ {
		jobChan <- worker.Job{ID: j, Name: fmt.Sprintf("Task-%d", j), Sleep: j}
	}

	select {} // block forever
}