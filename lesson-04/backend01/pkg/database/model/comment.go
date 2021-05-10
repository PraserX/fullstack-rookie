package model

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Comment   string    `json:"comment"`
	Timestamp time.Time `json:"timestamp"`
	UserID    uint      `json:"-"`
	User      User      `json:"user"`
}
