package users

import (
	"fmt"

	"github.com/sumudhar/go-book-store-user-api/datasources/mysql/users_db"
	"github.com/sumudhar/go-book-store-user-api/logger"
	"github.com/sumudhar/go-book-store-user-api/utils/errors"
	"github.com/sumudhar/go-book-store-user-api/utils/mysql_utils"
)

const(
	queryGetUser   			= "SELECT id,first_name,last_name, email, date_created,status from users WHERE ID= ?;"
	queryInsertUser			= "INSERT INTO  users (first_name, last_name, email,date_created,status,password) values (?, ?, ?, ?,?,?);"
	queryUpdateUser			= "UPDATE users set first_name= ? , last_name=?, email= ? WHERE ID= ?;"
	queryDeleteUser			= "DELETE FROM users where ID=?;"
	queryFindByUserStatus	= "SELECT id,first_name,last_name,email,date_created,status from users WHERE status=?;"
)
var (
	userDB= make(map[int64] *User)
)

func (user *User) Get() *errors.RestErr{
	stmt, getErr:= users_db.Client.Prepare(queryGetUser)
	if getErr != nil{
		logger.Error("error in while getting the user prepare statement",getErr)
        return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr:= result.Scan(&user.ID,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated,&user.Status); getErr!= nil {
		logger.Error("error when trying to scan user data in  get user by id",getErr)
		//return mysql_utils.ParseError(getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, saveErr:= users_db.Client.Prepare(queryInsertUser)
	if saveErr != nil{
		logger.Error("error in while saving the user prepare statement",saveErr)
        return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()	
	
	insertResult, saveErr := stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated,user.Status,user.Password)
	
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	} 
	userId, saveErr:= insertResult.LastInsertId()
	if saveErr != nil {
		logger.Error("error in saving user record",saveErr)
        return errors.NewInternalServerError("database error")
        //return mysql_utils.ParseError(saveErr)
	}
	user.ID= userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, updateErr:= users_db.Client.Prepare(queryUpdateUser)
	if updateErr != nil{
		logger.Error("error in while update the user prepare statement",updateErr)
        return errors.NewInternalServerError("database error")
        //return errors.NewInternalServerError(updateErr.Error())
	}
	defer stmt.Close()	
   
	_, updateErr = stmt.Exec(user.FirstName,user.LastName,user.Email,user.ID)
	if updateErr != nil {
		logger.Error("error in while update the user  ",updateErr)
        return errors.NewInternalServerError("database error")
		//return mysql_utils.ParseError(updateErr)
	} 
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, deleteErr:= users_db.Client.Prepare(queryDeleteUser)
	if deleteErr != nil{
		logger.Error("error in while deleting the user prepare statement",deleteErr)
        return errors.NewInternalServerError("database error")
        //return errors.NewInternalServerError(deleteErr.Error())
	}
	defer stmt.Close()
   	_, deleteErr = stmt.Exec(user.ID)
	if deleteErr != nil {
		logger.Error("error in while deleting the user ",deleteErr)
        return errors.NewInternalServerError("database error")
		//return mysql_utils.ParseError(deleteErr)
	} 
	return nil
}

func (user *User) FindByStatus(status string) ([] User, *errors.RestErr){
	 stmt,statusErr := users_db.Client.Prepare(queryFindByUserStatus)
	 
	 if statusErr !=nil{
		logger.Error("error in finding the user prepare statement",statusErr)
        return nil,errors.NewInternalServerError("database error")
		//return nil,errors.NewInternalServerError(statusErr.Error())
	 }
	 defer stmt.Close()

	 rows, statusErr := stmt.Query(status)
	 if statusErr !=nil{
        logger.Error("error in while finding user status ",statusErr)
        return nil, errors.NewInternalServerError("database error")
        //return nil, mysql_utils.ParseError(statusErr)
	 }
	 
	 defer rows.Close()
	 results := make([]User,0)
	 for rows.Next() {
		var user User;

		if statusErr := rows.Scan(&user.ID,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated,&user.Status); statusErr !=nil{
			logger.Error("error in while scannins the user object to user interface",statusErr)
            return nil, errors.NewInternalServerError("database error")
			//return nil, mysql_utils.ParseError(statusErr)
			
		}
       results = append(results,user)
	 }
     if len(results) ==0{
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching with status %s",status))
	 }
	 return results,nil

}