package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username	string	`json:"username" gorm:"uniqueIndex;not null"`
	Email		string	`json:"email" gorm:"uniqueIndex;not null"`
	Password	string	`json:"password" gorm:"not null"`
	Names		string	`json:"names"`
}

type NewUser struct {
	Username	string	`json:"username"`
	Email		string	`json:"email"`
}

type UpdateUserInput struct {
	Names	string	`json:"names"`
}

type UserPassInput struct {
	Password	string	`json:"password"`
}