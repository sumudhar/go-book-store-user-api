package users

import (
	"github.com/sumudhar/go-book-store-user-api/datasources/mysql/users_db"
	"github.com/sumudhar/go-book-store-user-api/utils/date_utils"
	"github.com/sumudhar/go-book-store-user-api/utils/errors"
	"github.com/sumudhar/go-book-store-user-api/utils/mysql_utils"
)

const (
	queryGetUser= "SELECT id,first_name,last_name, email, date_created from users WHERE ID= ?;"
	queryInsertUser = "INSERT INTO  users (first_name, last_name, email, date_created) values (?, ?, ?, ?);"
	queryUpdateUser= "UPDATE users set first_name= ? , last_name=?, email= ? WHERE ID= ?;"
	queryDeleteUser= "DELETE FROM users where ID=?;"
)
var (
	userDB= make(map[int64] *User)
)

func (user *User) Get() *errors.RestErr{
	stmt, getErr:= users_db.Client.Prepare(queryGetUser)
	if getErr != nil{
        return errors.NewInternalServerError(getErr.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr:= result.Scan(&user.ID,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated); getErr!= nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, saveErr:= users_db.Client.Prepare(queryInsertUser)
	if saveErr != nil{
        return errors.NewInternalServerError(saveErr.Error())
	}
	defer stmt.Close()	
	
	user.DateCreated= date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	} 
	userId, saveErr:= insertResult.LastInsertId()
	if saveErr != nil {
        return mysql_utils.ParseError(saveErr)
	}
	user.ID= userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, updateErr:= users_db.Client.Prepare(queryUpdateUser)
	if updateErr != nil{
        return errors.NewInternalServerError(updateErr.Error())
	}
	defer stmt.Close()	
   
	_, updateErr = stmt.Exec(user.FirstName,user.LastName,user.Email,user.ID)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	} 
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, deleteErr:= users_db.Client.Prepare(queryDeleteUser)
	if deleteErr != nil{
        return errors.NewInternalServerError(deleteErr.Error())
	}
	defer stmt.Close()

   	_, deleteErr = stmt.Exec(user.ID)
	if deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	} 
	return nil
}