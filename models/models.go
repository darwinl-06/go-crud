package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
	Year string `json:"year"`
}
