package model

type User struct {
	Id       uint   `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}