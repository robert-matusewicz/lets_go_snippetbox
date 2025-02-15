package mocks

import (
	"time"

	"snippetbox.matusewicz.dev/internal/models"
)

type UserModel struct{}

func (m *UserModel) Get(id int) (models.User, error) {
	if id == 1 {
		u := models.User{
			ID:      1,
			Name:    "Alice",
			Email:   "alice@example.com",
			Created: time.Now(),
		}
		return u, nil
	}
	return models.User{}, models.ErrNoRecord
}
func (m *UserModel) Insert(name string, email string, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email string, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) PasswordUpdate(id int, currentPassword string, newPassword string) error {
	if id == 1 && currentPassword == "pa$$word" {
		return nil
	}
	return models.ErrInvalidCredentials
}
