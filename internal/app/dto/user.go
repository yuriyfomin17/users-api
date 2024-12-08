package dto

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var (
	ErrRequired = errors.New("required value")
	ErrNegative = errors.New("negative value")
)

type User struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string
	Email     string
	Age       uint
	Created   time.Time
}

func (u User) Validate() error {
	if u.Firstname == "" {
		return fmt.Errorf("%w: first name", ErrRequired)
	}
	if u.Lastname == "" {
		return fmt.Errorf("%w: last name", ErrRequired)
	}
	if u.Email == "" {
		return fmt.Errorf("%w: email", ErrRequired)
	}
	if u.Age <= 0 {
		return fmt.Errorf("%w: age", ErrNegative)
	}
	return nil
}
func (u User) ValidateID() error {
	if u.ID == uuid.Nil {
		return fmt.Errorf("%w: id", ErrRequired)
	}
	return nil
}
