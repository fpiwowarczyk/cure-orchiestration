package worker

import (
	"fmt"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/fpiwowarczyk/cube-orchiestrator/task"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats")

}

func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("I will start task")
}

func (w *Worker) StopTask() {
	fmt.Println("I will stop task")
}
