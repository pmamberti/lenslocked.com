package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is returned when a resource cannot be found in
	// the database.
	ErrNotFound = errors.New("models: resource not found")
	// Err InvalidID returns when an invalid ID is provided.
	ErrInvalidID = errors.New("models: ID provided is invalid")
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

// first will query using the provided gorm.DB and it wwill get
// the first item returned, place it into dst. If nothing
// is found, it will return ErrNotFound.
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

// Close closes the UserService DB connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// DestructiveReset drops the user table and rebuilds it
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

// ByID will lookup the user by the ID provided
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil, otherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// ByEmail will lookup the user by the Email address provided
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil, otherError

func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// Create creates the provided user and the backfill data
// like ID, CreatedAt and UpdatedAt fields.
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// Update will update the proviced user with all the data
// in the provided user object.
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// Delete will delete the proviced user, selected by ID.
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}

	user := User{
		Model: gorm.Model{ID: id},
	}

	return us.db.Delete(&user).Error
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index`
}
