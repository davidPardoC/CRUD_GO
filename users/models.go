package users

import "encoding/json"

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email" gorm:"unique"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Created  int64  `gorm:"autoCreateTime"`
	Updated  int64  `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}

func (usr User) MarshalJSON() ([]byte, error) {
	var tmp struct {
		Email    string `json:"email"`
		Password string `json:"-"`
	}

	tmp.Email = usr.Email
	tmp.Password = usr.Password
	return json.Marshal(&tmp)
}
