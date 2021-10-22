package main

import (
	"sync"
	"errors"
)

type ToDo struct {
	lock *sync.RWMutex
	todos map[int]*UserToDo
}

func NewToDoStorage() *ToDo {
	return &ToDo{
		lock: &sync.RWMutex{},
		todos: make(map[int]*UserToDo),
	}
}

func (memStore *ToDo) Add(userId int, userToDo *UserToDo) error {
	memStore.lock.RLock()
	_, ok := memStore.todos[userId]
	memStore.lock.RUnlock()
	if ok {
		return errors.New(errAlreadyExistsUserToDo)
	}
	memStore.lock.Lock()
	memStore.todos[userId] = userToDo
	memStore.lock.Unlock()
	return nil
}

func (memStore *ToDo) Update(userId int, userToDo *UserToDo) error {
	memStore.lock.RLock()
	_, ok := memStore.todos[userId]
	memStore.lock.RUnlock()
	if !ok {
		return errors.New(errUserToDoDoesntExist)
	}

	memStore.lock.Lock()
	memStore.todos[userId] = userToDo
	memStore.lock.Unlock()
	return nil
}

func (memStore *ToDo) Delete(userId int) (*UserToDo, error) {
	memStore.lock.RLock()
	userToDo, ok := memStore.todos[userId]
	memStore.lock.RUnlock()
	if !ok {
		return &UserToDo{}, errors.New(errUserToDoDoesntExist)
	}

	delete(memStore.todos, userId)
	return userToDo, nil
}

func (memStore *ToDo) Get(userId int) (*UserToDo, error) {
	memStore.lock.RLock()
	userToDo, ok := memStore.todos[userId]
	memStore.lock.RUnlock()
	if !ok {
		return &UserToDo{}, errors.New(errUserToDoDoesntExist)
	}

	return userToDo, nil
}
