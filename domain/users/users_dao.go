package users

import (
	"fmt"
	"strings"

	"github.com/sumudhar/go-book-store-user-api/datasources/mysql/users_db"
	"github.com/sumudhar/go-book-store-user-api/utils/date_utils"
	"github.com/sumudhar/go-book-store-user-api/utils/errors"
)

const (
	indexUniqueEmail= "email_UNIQUE"
	errorNoRows= "no rows in result set"
	queryInsertUser = "INSERT INTO  users (first_name, last_name, email, date_created) values (?, ?, ?, ?);"
	queryGetUser= "SELECT id,first_name,last_name, email, date_created from users WHERE ID= ?;"
)
var (
	userDB= make(map[int64] *User)
)

func (user *User) Get() *errors.RestErr{

	stmt, err:= users_db.Client.Prepare(queryGetUser)
	if err != nil{
        return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID)
	if err:= result.Scan(&user.ID,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated); err!= nil {
		if strings.Contains(err.Error(),errorNoRows){
			return errors.NewBadRequestError(fmt.Sprintf("user not exists with the id: %d",user.ID))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get the user %d: %s",user.ID, err.Error()))
	}

	// if err := users_db.Client.Ping(); err !=nil{
	// 	panic(err)
	// }

	// result := userDB[user.ID]
	// if result ==nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("user %d not found",user.ID))
	// }
	// user.Email= result.Email
	// user.FirstName= result.FirstName
	// user.LastName= result.LastName
	// user.ID= result.ID
	// user.DateCreated= result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err:= users_db.Client.Prepare(queryInsertUser)
	if err != nil{
        return errors.NewInternalServerError(err.Error())

	}
	defer stmt.Close()
	
	user.DateCreated= date_utils.GetNowString()
	
	insertResult, err := stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated)
	
	if err != nil {
		if strings.Contains(err.Error(),indexUniqueEmail){
			return errors.NewBadRequestError(fmt.Sprintf("user email  %s  already exists",user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error occured while saving user",err.Error()))
	} 
	userId, err:= insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error to get last inserted id",err.Error()))  

	}
	user.ID= userId
	return nil
}