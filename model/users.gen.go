// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"mygpt/model_struct"
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID              string                   `gorm:"column:id;type:character varying(128);primaryKey" json:"id"`
	CreatedAt       *time.Time               `gorm:"column:created_at;type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	FirstName       *string                  `gorm:"column:first_name;type:text" json:"first_name"`
	LastName        *string                  `gorm:"column:last_name;type:text" json:"last_name"`
	Email           *string                  `gorm:"column:email;type:text" json:"email"`
	Status          *model_struct.UserStatus `gorm:"column:status;type:user_state" json:"status"`
	ProfileImageURL *string                  `gorm:"column:profile_image_url;type:text" json:"profile_image_url"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
