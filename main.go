package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
)

// Task holds the metadata about given task
type Task struct {
	ID           uuid.UUID
	IsCompleted  bool
	Status       string
	CreationTime time.Time
	TaskData     string
}

// getNewTask will return new task when it's get called
func getNewTask() *Task {
	return &Task{
		ID:           uuid.New(),
		IsCompleted:  false,
		Status:       "untouched",
		CreationTime: time.Now(),
		TaskData:     "I am new task",
	}
}

// getRandomStatus will retun random status everytime when it's get called
// added extra unwanted values to semulate the timeout
func getRandomStatus() string {
	rand.Seed(time.Now().UnixNano())
	s := []string{"completed", "failed", "retry", "networkDelay", "pending"}
	return s[rand.Intn(len(s))]
}

// TODO: try to implement the taskAdder func
// func taskAdder(){
//
// }

// taskExecutor will execute the task and assign the random status when it's done
func taskExecutor(task *Task) *Task {
	// m := sync.Mutex{}
	// waiting for random time to semulate the actual work or network delay
	rand.Seed(time.Now().UnixNano())
	t := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(t)

	// m.Lock()
	// log.Println(task)
	log.Println("executing task...: ", task.ID)
	status := getRandomStatus()
	// log.Println("random status: ", status)

	// check if task is completed
	// if we get randomStatus as 'completed' change the task field
	if status == "completed" {
		task.IsCompleted = true
		// m.Unlock()
		log.Println("task completed: ", task.ID)
		return nil
	}
	return task
}

func main() {

	// NOTE: code can be refactored and written in diff functions.
	// For simplicity I have created all gorutines in main

	// wg := sync.WaitGroup{}
	// No of task
	n := 50

	// defined the queue type chan
	queue := make(chan *Task, n)
	// failedQueue := make(chan *Task, n)

	// Create N number of task with random time intervals
	go func() {
		for i := 1; i <= n; i++ {
			t := time.Duration(rand.Intn(200)) * time.Millisecond
			time.Sleep(t)

			tsk := getNewTask()
			// log.Println("creating task and put it into queue")
			queue <- tsk
		}
	}()

	// wg.Add(2)
	// task executoer gorutine
	go func() {
		for t := range queue {
			tsk := taskExecutor(t)
			// check if task is completed (nil) or failed
			// if failed put back to the queue
			if tsk != nil {
				queue <- tsk
				// failedQueue <- tsk
			}
		}
		// wg.Done()
	}()

	// cleanup task
	go func() {
		// get the task from the queue and check if it's completed
		// t := <-queue
		for t := range queue {
			// get the time taken by the task
			// NOTE: timeout can be handled using context with timeout
			remaningTime := time.Since(t.CreationTime)
			// log.Println(remaningTime)

			//cancel the task if creation time is more than 500 milsec
			// NOTE: hardcoded timeout can be used as var
			if remaningTime < 500*time.Millisecond {
				// t := <-failedQueue
				if !t.IsCompleted {
					// log.Println("taks is not completed puting back to the queue")
					queue <- t
				}
			} else {
				log.Println("Timeout for the task...", t.ID)
			}
			// log.Println("is task completed: ", t.IsCompleted)
		}
		// wg.Done()

	}()

	// wg.Wait()
	// chan to listen for OS level interuption
	sigChan := make(chan os.Signal, 1)

	// register signal to catch on sigChan chan
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// wait until it's iterupted by OS
	fmt.Println("Service exited due to: ", <-sigChan)

}
