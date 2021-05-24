package controllers

import (
	"net/http"

	"lenslocked.com/views"
)

// NewUsers is used to create a new Users controller
// It will panic if the templtes are not parsed correctly
// and should only be used during initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}
