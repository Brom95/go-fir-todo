package main

import "github.com/google/uuid"

type TaskStorage interface {
	Add(Task) error
	Get(uuid.UUID) *Task
	UpdateStatus(uuid.UUID, bool) error
	GetAll() []*Task
}
type Task struct {
	Complete bool
	Name     string
	Id       uuid.UUID
}
type TaskIdResponse struct {
	TaskId string
}
