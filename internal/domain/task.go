package domain

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Priority    int    `json:"priority" gorm:"column:priority"`
}

type TaskRequest struct {
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Priority    int    `json:"priority" gorm:"column:priority"`
}
