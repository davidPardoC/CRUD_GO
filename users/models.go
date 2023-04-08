package users

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email" gorm:"unique"  binding:"required"`
	Password string `json:"password,omitempty"  binding:"required"`
	Created  int64  `gorm:"autoCreateTime"`
	Updated  int64  `gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
