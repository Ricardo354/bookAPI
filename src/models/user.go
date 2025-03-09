package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" json:"usuario"`
	Senha string `json:"senha"`
}