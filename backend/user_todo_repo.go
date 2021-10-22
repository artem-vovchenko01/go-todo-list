package main

import (
	"errors"
	"sync"
)

const alreadyToDoListExistsErr = "ToDoList by such ID already exist"
const toDoListDoesntExistErr = "There is no ToDoList by such ID"

type UserToDo struct {
	currentId int
	lock  *sync.RWMutex
	lists map[int]*ToDoList
}

func NewUserToDoStorage() *UserToDo {
	return &UserToDo{
		currentId: 1,
		lock:  &sync.RWMutex{},
		lists: make(map[int]*ToDoList),
	}
}

func (memStore *UserToDo) Add(toDoList *ToDoList) error {
	memStore.lock.RLock()
	_, ok := memStore.lists[memStore.currentId]
	memStore.lock.RUnlock()
	if ok {
		return errors.New(alreadyToDoListExistsErr)
	}
	memStore.lock.Lock()
	toDoList.Id = memStore.currentId
	memStore.lists[memStore.currentId] = toDoList
	memStore.currentId += 1
	memStore.lock.Unlock()
	return nil
}

func (memStore *UserToDo) Update(listId int, toDoList *ToDoList) error {
	memStore.lock.RLock()
	_, ok := memStore.lists[listId]
	memStore.lock.RUnlock()
	if !ok {
		return errors.New(toDoListDoesntExistErr)
	}

	memStore.lock.Lock()
	memStore.lists[listId] = toDoList
	memStore.lock.Unlock()
	return nil
}

func (memStore *UserToDo) Delete(listId int) (*ToDoList, error) {
	memStore.lock.RLock()
	toDoList, ok := memStore.lists[listId]
	memStore.lock.RUnlock()
	if !ok {
		return &ToDoList{}, errors.New(toDoListDoesntExistErr)
	}

	delete(memStore.lists, listId)
	return toDoList, nil
}

func (memStore *UserToDo) Get(listId int) (*ToDoList, error) {
	memStore.lock.RLock()
	toDoList, ok := memStore.lists[listId]
	memStore.lock.RUnlock()
	if !ok {
		return &ToDoList{}, errors.New(toDoListDoesntExistErr)
	}

	return toDoList, nil
}
