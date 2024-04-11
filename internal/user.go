package just_auth

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
}
