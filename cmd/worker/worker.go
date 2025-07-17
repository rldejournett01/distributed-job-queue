package worker

import (
	"fmt"
	"time"
)

//Job represents a basic job struct
type Job struct{
	ID int
	Name string
	Sleep int

}

func Worker(id int, jobChan <-chan Job) {
	for job := range jobChan {
		fmt.Printf("[Worker %d] Starting job: %s (sleep %d)\n", id, job.Name, job.Sleep)
		time.Sleep(time.Duration(job.Sleep) * time.Second)
		fmt.Printf("[Worker %d] Finished job: %s\n", id, job.Name)
	}
}