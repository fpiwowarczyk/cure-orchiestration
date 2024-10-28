package manager

import (
	"cube-orchiestrator/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.Task
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropirate wokrer")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("I will updat tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}
