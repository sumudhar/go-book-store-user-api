package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sumudhar/go-book-store-user-api/domain/users"
	"github.com/sumudhar/go-book-store-user-api/services"
	"github.com/sumudhar/go-book-store-user-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
    // fmt.Println(user)
	// bytes,err := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(bytes))
	// fmt.Println(err)
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
	
	//c.String(http.StatusNotImplemented, "Not yet Implemented!!!")

}

func GetUser(c *gin.Context) {

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

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "Not yet Implemented!!!")
// }
