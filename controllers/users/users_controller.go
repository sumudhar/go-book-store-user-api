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

	user, getErr:= services.UsersService.GetUser(userId) 

	if getErr !=nil {
		c.JSON(getErr.Status,getErr)
        return
	}

	c.JSON(http.StatusOK,user.Marshall(c.GetHeader("X-Public") == "true"))

	
}

func Create(c *gin.Context) {

	var user users.User

	// c.ShouldBindJSON() this method will capture the object come from request and bind the values
	// againist the User Model columns and verifies

	if err:= c.ShouldBindJSON(&user); err!=nil {
		restErr := errors.NewBadRequestError("invalid Json body") 
		c.JSON(restErr.Status,restErr)		
		return
	}
	
	result, saveErr:= services.UsersService.CreateUser(user) 
	if saveErr !=nil {
		// Implement the error
		c.JSON(saveErr.Status,saveErr)
        return
	}
	
	c.JSON(http.StatusCreated,result.Marshall(c.GetHeader("X-Public") == "true"))
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
	
	result, updateErr:= services.UsersService.UpdateUser(isPartial,user) 
	
	if updateErr !=nil {
		// Implement the error
		c.JSON(updateErr.Status,updateErr)
        return
	}
	c.JSON(http.StatusOK,result.Marshall(c.GetHeader("X-Public") == "true"))
}


func Delete(c *gin.Context){
	userId, idErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if idErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status,err)		
		return
	}
	deleteErr:= services.UsersService.DeleteUser(userId) 
	if deleteErr !=nil {
		// Implement the error
		c.JSON(deleteErr.Status,deleteErr)
        return
	}
	c.JSON(http.StatusOK,map[string]string{"status": "deleted"} )
}

func Search( c *gin.Context){
	status := c.Query("status")
	users, statusErr := services.UsersService.SearchUser(status)
	
	if statusErr != nil{
		c.JSON(statusErr.Status,statusErr)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))

}