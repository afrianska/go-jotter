package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Tittle		string	`json:"title" gorm:"not null"`
	Description	string	`json:"description" gorm:"not null"`
	Amount		int	`json:"amount" gorm:"not null"`
}

type UpdateProductInput struct {

	Tittle		string	`json:"title"`
	Description	string	`json:"description"`
	Amount		int	`json:"amount" default:"0"`
}