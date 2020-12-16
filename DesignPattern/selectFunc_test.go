package DesignPattern

import (
	"testing"
)

func TestSelectFunc(t *testing.T) {
	task := NewTask(WithFamilyName("Parker"), WithFirstName("Peter"))
	t.Log(task.String())
}

type Task struct {
	FirstName  string
	FamilyName string
}

type Option func(task *Task)

func WithFamilyName(s string) Option {
	return func(task *Task) {
		task.FamilyName = s
	}
}
func WithFirstName(s string) Option {
	return func(task *Task) {
		task.FirstName = s
	}
}

func NewTask(opts ...Option) *Task {
	task := &Task{}
	for _, value := range opts {
		value(task)
	}
	return task
}

func (t *Task) String() string {
	return t.FirstName + "." + t.FamilyName
}
