package users

type UserRepository interface {
	InsertUser(user UserModel) UserModel
	GetAllUsers() []UserModel
	GetUserById(id int) UserModel
}

var UserRepositoryImplementation UserRepository

func SetUserRepositoryImplementation(repository UserRepository) {
	UserRepositoryImplementation = repository
}
