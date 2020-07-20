package users

import (
	"strings"

	"github.com/sumudhar/go-book-store-user-api/utils/errors"
)

const (

	ActiveStatus= "active"
)
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`

}

func (user *User) Validate() *errors.RestErr {
	
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	
	if user.Email == "" {
		return  errors.NewBadRequestError("Invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)

	if user.Password == "" {
		return  errors.NewBadRequestError("Invalid password")
	}
	return nil

}

