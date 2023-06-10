package models

import "gorm.io/gorm"

type Product struct {
	*gorm.Model

	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int32  `db:"price"`
	Image       string `db:"image"`
	Category    string `db:"category"`
}
