package main

import (
	"errors"
	"fmt"
	"sync"
)

const errAlreadyUserExists = "Task by such ID already exist"
const errUserDoesntExist = "There is no task by such ID"

type User struct {
	id       int
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStorage struct {
	currentId int
	lock      *sync.RWMutex
	users     map[int]User
}

type UserIdStorage struct {
	lock    sync.RWMutex
	userIds map[string]int
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		currentId: 1,
		lock:      &sync.RWMutex{},
		users:     make(map[int]User),
	}
}

func NewUserIdStorage() *UserIdStorage {
	return &UserIdStorage{
		lock:    sync.RWMutex{},
		userIds: make(map[string]int),
	}
}

func (memStore *UserStorage) Add(user User) error {
	fmt.Println("user adding")
	fmt.Println(user)
	memStore.lock.RLock()
	_, ok := memStore.users[memStore.currentId]
	memStore.lock.RUnlock()
	if ok {
		return errors.New(errAlreadyTaskExists)
	}
	user.id = memStore.currentId
	memStore.lock.Lock()
	memStore.users[memStore.currentId] = user
	memStore.currentId += 1
	memStore.lock.Unlock()
	Storage.userIds.Add(user.Email, user.id)
	return nil
}

func (memStore *UserStorage) Update(userId int, user User) error {
	memStore.lock.RLock()
	_, ok := memStore.users[userId]
	memStore.lock.RUnlock()
	if !ok {
		return errors.New(errTaskDoesntExist)
	}

	memStore.lock.Lock()
	memStore.users[userId] = user
	memStore.lock.Unlock()
	return nil
}

func (memStore *UserStorage) Delete(userId int) (User, error) {
	memStore.lock.RLock()
	user, ok := memStore.users[userId]
	memStore.lock.RUnlock()
	if !ok {
		return User{}, errors.New(errTaskDoesntExist)
	}

	delete(memStore.users, userId)
	return user, nil
}

func (memStore *UserStorage) Get(userId int) (User, error) {
	memStore.lock.RLock()
	task, ok := memStore.users[userId]
	memStore.lock.RUnlock()
	if !ok {
		return User{}, errors.New(errTaskDoesntExist)
	}

	return task, nil
}

func (memStore *UserIdStorage) Add(email string, userId int) error {
	memStore.lock.RLock()
	_, ok := memStore.userIds[email]
	memStore.lock.RUnlock()
	if ok {
		return errors.New(errAlreadyTaskExists)
	}
	memStore.lock.Lock()
	memStore.userIds[email] = userId
	memStore.lock.Unlock()
	return nil
}

func (memStore *UserIdStorage) Update(email string, userId int) error {
	memStore.lock.RLock()
	_, ok := memStore.userIds[email]
	memStore.lock.RUnlock()
	if !ok {
		return errors.New(errTaskDoesntExist)
	}

	memStore.lock.Lock()
	memStore.userIds[email] = userId
	memStore.lock.Unlock()
	return nil
}

func (memStore *UserIdStorage) Delete(email string) (int, error) {
	memStore.lock.RLock()
	userId, ok := memStore.userIds[email]
	memStore.lock.RUnlock()
	if !ok {
		return -1, errors.New(errTaskDoesntExist)
	}

	delete(memStore.userIds, email)
	return userId, nil
}

func (memStore *UserIdStorage) Get(email string) (int, error) {
	memStore.lock.RLock()
	userId, ok := memStore.userIds[email]
	memStore.lock.RUnlock()
	if !ok {
		return -1, errors.New(errTaskDoesntExist)
	}

	return userId, nil
}
