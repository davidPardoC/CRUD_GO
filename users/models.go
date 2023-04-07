package users

type UserModel struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Created  int64  `gorm:"autoCreateTime"`
	Updated  int64  `gorm:"autoUpdateTime"`
}
