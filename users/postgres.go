package users

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db gorm.DB
}

func (repo UserRepo) InsertUser(user User) (User, error) {
	result := repo.Db.Create(&user)
	err := result.Error
	if err != nil {
		return user, errors.New(user.Email + " " + err.Error())
	}
	return user, nil
}

func (repo UserRepo) GetAllUsers() []User {
	var users []User
	repo.Db.Select("email", "id").Find(&users)
	return users
}

func (repo UserRepo) GetUserById(id int) User {
	var user User
	repo.Db.First(&user, id)
	return user
}

func (repo UserRepo) GetUserByEmail(email string) User {
	var user User
	repo.Db.Where("name = ?", email).First(&user)
	return user
}
