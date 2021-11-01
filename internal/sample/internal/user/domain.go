package user

import (
	"strings"
	"time"

	"github.com/arimaulana/point-of-no-return/internal/common/pkg/generator"
)

// User represents a user domain
type User struct {
	ID        string `json:"id"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Password  string
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewUser initiates user
func NewUser(username, email, password string) User {
	u := User{
		ID:        generator.GenerateUUID(),
		UserName:  username,
		Email:     strings.ToLower(email),
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return u
}

// GetID get user id
func (u User) GetID() string {
	return u.ID
}

// GetUserName get user name
func (u User) GetUserName() string {
	return u.UserName
}

// GetEmail get email
func (u User) GetEmail() string {
	return u.Email
}
