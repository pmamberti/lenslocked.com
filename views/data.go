package views

import (
	"log"

	"lenslocked.com/models"
)

const (
	AlertLvlError   = "danger"
	AlertLvlWarning = "warning"
	AlertLvlInfo    = "info"
	AlertLvlSuccess = "success"

	// AlertMsgGeneric is displayed whenever any random error is
	// encountered by our backend.
	AlertMsgGeneric = "Something went wrong. Please try again, and contact us if the problem persists."
)

// Alert is used to render bootstrap alert messages in the
// templates
type Alert struct {
	Level   string
	Message string
}

// Data is the top level struct that views expect data to
// come in.
type Data struct {
	Alert *Alert
	User  *models.User
	Yield interface{}
}

func (d *Data) SetAlert(err error) {
	if pErr, ok := err.(PublicError); ok {
		d.Alert = &Alert{
			Level:   AlertLvlError,
			Message: pErr.Public(),
		}
	} else {
		d.Alert = &Alert{
			Level:   AlertLvlError,
			Message: AlertMsgGeneric,
		}
	}
	log.Println(err)
}

func (d *Data) AlertError(msg string) {
	d.Alert = &Alert{
		Level:   AlertLvlError,
		Message: msg,
	}
}

type PublicError interface {
	error
	Public() string
}
