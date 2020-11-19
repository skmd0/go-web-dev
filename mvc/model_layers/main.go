package mail

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

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

// This can optionally be an interface with constructor that initializes it's struct
type UserService struct {
	UserDB
}

// Method on top layer specific to UserService that is not CRUD db method
func (us *UserService) Authenticate(user *User) error {
	return nil
}

type userValidator struct {
	UserDB
}

// Validator level method that check if conditions are met and then delegates work forward
func (uv *userValidator) Create(user *User) error {
	if user.ID <= 0 {
		return errors.New("ID cannot be lower then 1")
	}
	return uv.UserDB.Create(user)
}

var _ UserDB = &userGorm{}

type userGorm struct {
	db *gorm.DB
}

func (ug *userGorm) ByID(id uint) (*User, error) {
	return nil, nil
}

func (ug *userGorm) Create(user *User) error {
	return nil
}

func (ug *userGorm) Update(user *User) error {
	return nil
}

func (ug *userGorm) Delete(id uint) error {
	return nil
}

func (ug *userGorm) AutoMigrate() error {
	return nil
}

func (ug *userGorm) DestructiveReset() error {
	return nil
}

func NewUserService() *UserService {
	var db *gorm.DB
	// init the DB connection
	return &UserService{
		UserDB: &userValidator{
			UserDB: &userGorm{
				db: db,
			},
		},
	}
}

func main() {
	us := NewUserService()
	fmt.Printf("%+v\n", us)
}
