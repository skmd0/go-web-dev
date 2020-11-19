package mail

import (
	"errors"
	"gorm.io/gorm"
)

type userValFunc func(user *User) error

// this is a helper function which will loop through all the validators/normalizers
func runUserValFunc(user *User, fns ...userValFunc) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

type userValidator struct {
	UserDB
}

func (uv *userValidator) Create(user *User) error {
	// here specify which validators to run on the provided User
	if err := runUserValFunc(user, uv.checkEmail, uv.checkPassword); err != nil {
		return err
	}
	return uv.UserDB.Create(user)
}

func (uv *userValidator) checkEmail(user *User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (uv *userValidator) checkPassword(user *User) error {
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

// the outer service and inner gorm layer are missing from this example, look at model_layers example for that

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;uniqueIndex"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;uniqueIndex"`
}

type UserDB interface {
	ByID(id uint) (*User, error)

	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	AutoMigrate() error
	DestructiveReset() error
}
