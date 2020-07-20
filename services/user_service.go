package services

import (
	"github.com/sumudhar/go-book-store-user-api/domain/users"
	"github.com/sumudhar/go-book-store-user-api/utils/crypto_utils"
	"github.com/sumudhar/go-book-store-user-api/utils/date_utils"
	"github.com/sumudhar/go-book-store-user-api/utils/errors"
)
var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User,* errors.RestErr)
	CreateUser(users.User) (*users.User, * errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, * errors.RestErr)
	DeleteUser(int64) * errors.RestErr
	SearchUser(string) (users.Users, * errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {

	result := &users.User{ID: userId}

	if err := result.Get(); err != nil {
       return nil, err
	}
	return  result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
       return nil, err
	}
	user.DateCreated= date_utils.GetNowDBFormat()
	user.Status= users.ActiveStatus
	user.Password=crypto_utils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return  &user, nil

}

func (s *usersService) UpdateUser(isPartial bool ,user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{ID: user.ID}

	if err := current.Get(); err != nil {
		return nil, err
	}

	if isPartial{
		
		if user.Email !=""{
			current.Email= user.Email
		}
		if user.FirstName !=""{
			current.FirstName= user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}

	}else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
    
	if err := current.Update(); err != nil {
		return nil, err
	}
	return  current, nil

}



func (s *usersService) DeleteUser(userID int64)  *errors.RestErr{

	user := &users.User{ID: userID}
	
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr){
    dao := &users.User{}
	return dao.FindByStatus(status)
	   	
}