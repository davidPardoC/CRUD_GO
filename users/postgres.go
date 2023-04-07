package users

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error on vconecting database", err.Error())
		return nil, err
	}
	return db, nil
}

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
