package users

type UserRepository interface {
	InsertUser(user User) (User, error)
	GetAllUsers() []User
	GetUserById(id int) User
	GetUserByEmail(email string) User
}

var UserRepositoryImplementation UserRepository

func SetUserRepositoryImplementation(repository UserRepository) {
	UserRepositoryImplementation = repository
}
