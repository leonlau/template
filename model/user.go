package model

import "template/pkg/auth"

type UserModel struct {
	BaseModel
	Username string `from:"username" json:"username" gorm:"column:username;not null"`
	Password string `from:"password" json:"password" gorm:"column:password;not null"`
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}
