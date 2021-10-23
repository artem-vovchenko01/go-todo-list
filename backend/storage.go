package main

func NewStorageService() *StorageService {
	return &StorageService{
		toDos:   NewToDoStorage(),
		users:   NewUserStorage(),
		userIds: NewUserIdStorage(),
	}
}

type StorageService struct {
	toDos   *ToDo
	users   *UserStorage
	userIds *UserIdStorage
}

func (s *StorageService) GetUserToDoByEmail(email string) (*UserToDo, error) {
	userId, err := Storage.userIds.Get(email)
	if err != nil {
		return &UserToDo{}, err
	}
	lists, err := Storage.toDos.Get(userId)
	if err != nil {
		return &UserToDo{}, err
	}
	return lists, nil
}

func (s *StorageService) GetOrCreateUserToDoByEmail(email string) (*UserToDo, error) {
	userId, err := Storage.userIds.Get(email)
	if err != nil {
		return &UserToDo{}, err
	}
	lists, err := Storage.toDos.Get(userId)
	if err != nil {
		Storage.toDos.Add(userId, NewUserToDoStorage())
		lists, _ = Storage.toDos.Get(userId)
	}
	return lists, nil
}
