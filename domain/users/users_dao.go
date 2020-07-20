package users

import (
	"fmt"

	"github.com/sumudhar/go-book-store-user-api/datasources/mysql/users_db"
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
        return errors.NewInternalServerError(getErr.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr:= result.Scan(&user.ID,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated,&user.Status); getErr!= nil {
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
	
	insertResult, saveErr := stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated,user.Status,user.Password)
	
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

func (user *User) FindByStatus(status string) ([] User, *errors.RestErr){
	 stmt,statusErr := users_db.Client.Prepare(queryFindByUserStatus)
	 
	 if statusErr !=nil{
		 return nil,errors.NewInternalServerError(statusErr.Error())
	 }
	 defer stmt.Close()

	 rows, statusErr := stmt.Query(status)
	 if statusErr !=nil{
        return nil, mysql_utils.ParseError(statusErr)
	 }
	 
	 defer rows.Close()
	 fmt.Println("I am ok  1, ?",rows)
	 results := make([]User,0)
	 for rows.Next() {
		fmt.Println("I am ok  2")
		var user User;

		if statusErr := rows.Scan(&user.ID,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated,&user.Status); statusErr !=nil{
			fmt.Println("I am not ok  3, ?",user)
			return nil, mysql_utils.ParseError(statusErr)
			
		}
       results = append(results,user)
	 }
     fmt.Println("I am ok  3")
	 if len(results) ==0{
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching with status %s",status))
	 }
	 return results,nil

}