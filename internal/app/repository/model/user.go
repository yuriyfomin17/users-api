package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `pg:"id"`
	FirstName string    `pg:"first_name"`
	LastName  string    `pg:"last_name"`
	Email     string    `pg:"email"`
	Age       uint      `pg:"age"`
	CreatedAt time.Time `pg:"created_at,type:timestamp(6)"`
}
