package users

import (
	"fmt"

	"github.com/sumudhar/go-book-store-user-api/utils/errors"
)


var (
	userDB= make(map[int64] *User)
)

func (user *User) Get() *errors.RestErr{
	result := userDB[user.ID]
	if result ==nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found",user.ID))
	}
	user.Email= result.Email
	user.FirstName= result.FirstName
	user.LastName= result.LastName
	user.ID= result.ID
	user.DateCreated= result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	var current = userDB[user.ID] 
    if current != nil {
		if current.Email == user.Email{
			return errors.NewBadRequestError(fmt.Sprintf("user email  %s  already exists",user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user id  %d already exists",user.ID))
	}
	userDB[user.ID]= user

	return nil
}