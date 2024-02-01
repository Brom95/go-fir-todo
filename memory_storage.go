package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type MemoryStorage struct {
	mutex *sync.RWMutex
	data  map[uuid.UUID]*Task
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		mutex: &sync.RWMutex{},
		data:  make(map[uuid.UUID]*Task),
	}
}

// UpdateStatus implements TaskStorage.
func (ms *MemoryStorage) UpdateStatus(id uuid.UUID, complete bool) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	if task, exists := ms.data[id]; exists {
		task.Complete = complete
		return nil
	}
	return fmt.Errorf("task with ID %v not found", id)
}

// Add implements TaskStorage.
func (ms *MemoryStorage) Add(task Task) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	task.Id = uuid.New()
	ms.data[task.Id] = &task
	return nil
}

// Get implements TaskStorage.
func (ms *MemoryStorage) Get(id uuid.UUID) *Task {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()
	if task, exists := ms.data[id]; exists {
		return task
	}
	return &Task{}
}

// GetAll implements TaskStorage.
func (ms *MemoryStorage) GetAll() []*Task {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()
	var tasks []*Task
	for _, task := range ms.data {
		tasks = append(tasks, task)
	}
	return tasks
}
