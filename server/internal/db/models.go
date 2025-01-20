package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID		uuid.UUID	`gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username	string		`gorm:"unique;not null"`
	CreatedAt	time.Time	`gorm:"autoCreateTime"`
	UpdatedAt	time.Time	`gorm:"autoUpdateTime"`
	DeletedAt	gorm.DeletedAt	`gorm:"index"`
}
