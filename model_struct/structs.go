package model_struct

import "database/sql/driver"

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
)

func (s *UserStatus) Scan(value interface{}) error {
	*s = UserStatus(value.([]byte))
	return nil
}

func (s UserStatus) Value() (driver.Value, error) {
	return string(s), nil
}
