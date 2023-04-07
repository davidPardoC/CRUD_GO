package users

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db gorm.DB
}

func (repo UserRepo) InsertUser(user UserModel) UserModel {
	repo.Db.Create(&user)
	fmt.Println(user.Id)
	return user
}

func (repo UserRepo) GetAllUsers() []UserModel {
	var users []UserModel
	repo.Db.Find(&users)
	return users
}

func (repo UserRepo) GetUserById(id int) UserModel {
	var user UserModel
	repo.Db.First(&user, id)
	return user
}
