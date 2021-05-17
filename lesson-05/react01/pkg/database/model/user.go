package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}
