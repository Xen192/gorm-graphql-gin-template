package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ParentId  *string
	FileName  *string
	MimeType  string
	Data      []byte
}
