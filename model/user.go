package model

type LoginParams struct {
	UserName string
	Password string
}

type User struct {
	ID       int32  `json:"id" gorm:"column:id"`
	UserName string `json:"user_name" gorm:"column:user_name"`
	PassWord string `json:"password" gorm:"column:password"`
}
