package models

import (
	"fmt"
	"testing"
	"time"
)

func testingUserService() (*UserService, error) {
	const (
		host    = "localhost"
		port    = 5432
		user    = "piero"
		db_name = "lenslocked_test"
	)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		db_name,
	)

	us, err := NewUserService(dsn)
	if err != nil {
		return nil, err
	}
	us.db.LogMode(false)
	// Clear the users table between tests
	us.DestructiveReset()
	return us, nil
}

func TestCreateUser(t *testing.T) {
	us, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}

	user := User{ //
		Name:  "Michael Scott",
		Email: "michael@example.com",
	}
	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID == 0 {
		t.Errorf("expected ID greater than 0 received %d", user.ID)
	}

	if time.Since(user.CreatedAt) > time.Duration(5*time.Second) {
		t.Errorf(
			"expected created_at to be less than 5s ago - received %v",
			user.CreatedAt,
		)
	}
}
