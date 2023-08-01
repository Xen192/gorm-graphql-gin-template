package models

import (
	"time"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

type User struct {
	ID              string `gorm:"primaryKey"`
	CreatedAt       time.Time
	FirstName       *string
	LastName        *string
	Email           *string
	Status          Status `gorm:"type:text"`
	ProfileImageURL *string
}

type ClerkUser struct {
	ID             string `gorm:"primaryKey"`
	LinkedIdentity string
}
