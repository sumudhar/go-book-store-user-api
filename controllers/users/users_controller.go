package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sumudhar/go-book-store-user-api/domain/users"
	"github.com/sumudhar/go-book-store-user-api/services"
	"github.com/sumudhar/go-book-store-user-api/utils/errors"
)

func Get(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)

	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status,err)		
		return
	}

	user, getErr:= services.GetUser(userId) 

	if getErr !=nil {
		// Implement the error
		c.JSON(getErr.Status,getErr)
        return
	}

	c.JSON(http.StatusOK,user)

	
}

func Create(c *gin.Context) {

	var user users.User

	if err:= c.ShouldBindJSON(&user); err!=nil {
		restErr := errors.NewBadRequestError("invalid Json body") 
		c.JSON(restErr.Status,restErr)		
		return
	}
	result, saveErr:= services.CreateUser(user) 
	if saveErr !=nil {
		// Implement the error
		c.JSON(saveErr.Status,saveErr)
        return
	}
	c.JSON(http.StatusCreated,result)
}

func Update(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status,err)		
		return
	}
	
	var user users.User
	if err:= c.ShouldBindJSON(&user); err!=nil {
		restErr := errors.NewBadRequestError("invalid Json body") 
		c.JSON(restErr.Status,restErr)		
		return
	}

	user.ID = userId
	isPartial := c.Request.Method == http.MethodPatch
	
	result, updateErr:= services.UpdateUser(isPartial,user) 
	
	if updateErr !=nil {
		// Implement the error
		c.JSON(updateErr.Status,updateErr)
        return
	}
	c.JSON(http.StatusOK,result)
}


func Delete(c *gin.Context){
	userId, idErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if idErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status,err)		
		return
	}
	deleteErr:= services.DeleteUser(userId) 
	if deleteErr !=nil {
		// Implement the error
		c.JSON(deleteErr.Status,deleteErr)
        return
	}
	c.JSON(http.StatusOK,map[string]string{"status": "deleted"} )
}

func Search( c *gin.Context){
	status := c.Query("status")
	users, statusErr := services.Search(status)
	
	if statusErr != nil{
		c.JSON(statusErr.Status,statusErr)
		return
	}
	c.JSON(http.StatusOK,users)

}