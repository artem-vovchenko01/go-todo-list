package main

import (
	"errors"
	"sync"
)

type ToDoList struct {
	currentTaskId int
	lock          *sync.RWMutex
	Name          string `json:"name"`
	Id            int    `json:"id"`
	tasks         map[int]*Task
}

func NewToDoListStorage() *ToDoList {
	return &ToDoList{
		currentTaskId: 1,
		lock:          &sync.RWMutex{},
		tasks:         make(map[int]*Task),
	}
}

func (memStore *ToDoList) Add(task *Task) error {
	memStore.lock.RLock()
	_, ok := memStore.tasks[memStore.currentTaskId]
	memStore.lock.RUnlock()
	if ok {
		return errors.New(errAlreadyTaskExists)
	}
	memStore.lock.Lock()
	task.Id = memStore.currentTaskId
	task.Status = "open"
	memStore.tasks[memStore.currentTaskId] = task
	memStore.currentTaskId += 1
	memStore.lock.Unlock()
	return nil
}

func (memStore *ToDoList) Update(taskId int, task *Task) error {
	memStore.lock.RLock()
	_, ok := memStore.tasks[taskId]
	memStore.lock.RUnlock()
	if !ok {
		return errors.New(errTaskDoesntExist)
	}

	memStore.lock.Lock()
	memStore.tasks[taskId] = task
	memStore.lock.Unlock()
	return nil
}

func (memStore *ToDoList) Delete(taskId int) (*Task, error) {
	memStore.lock.RLock()
	task, ok := memStore.tasks[taskId]
	memStore.lock.RUnlock()
	if !ok {
		return &Task{}, errors.New(errTaskDoesntExist)
	}

	delete(memStore.tasks, taskId)
	return task, nil
}

func (memStore *ToDoList) Get(taskId int) (*Task, error) {
	memStore.lock.RLock()
	task, ok := memStore.tasks[taskId]
	memStore.lock.RUnlock()
	if !ok {
		return &Task{}, errors.New(errTaskDoesntExist)
	}

	return task, nil
}
