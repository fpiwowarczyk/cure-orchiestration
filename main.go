package main

import (
	"fmt"
	"time"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/fpiwowarczyk/cube-orchiestrator/manager"
	"github.com/fpiwowarczyk/cube-orchiestrator/node"
	"github.com/fpiwowarczyk/cube-orchiestrator/task"
	"github.com/fpiwowarczyk/cube-orchiestrator/worker"
	"github.com/google/uuid"
)

func main() {
	t := task.Task{
		ID:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "Image-1",
		Memory: 1024,
		Disk:   1,
	}

	te := task.TaskEvent{
		ID:        uuid.New(),
		State:     task.Pending,
		Timestamp: time.Now(),
		Task:      t,
	}

	fmt.Printf("task: %v\n", t)
	fmt.Printf("taskEvent: %v\n", te)

	w := worker.Worker{
		Name:  "worker-1",
		Queue: *queue.New(10),
		Db:    make(map[uuid.UUID]*task.Task),
	}

	fmt.Printf("worker: %v\n", w)
	w.CollectStats()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	m := manager.Manager{
		Pending: *queue.New(10),
		TaskDb:  make(map[string][]task.Task),
		EventDb: make(map[string][]task.TaskEvent),
		Workers: []string{w.Name},
	}

	fmt.Printf("manager: %v\n", m)
	m.SelectWorker()
	m.UpdateTasks()
	m.SendWork()

	n := node.Node{
		Name:   "node-1",
		Ip:     "192.169.1.1",
		Cores:  4,
		Memory: 1024,
		Disk:   25,
		Role:   "Worker",
	}

	fmt.Printf("node: %v\n", n)
}
