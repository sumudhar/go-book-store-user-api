package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sumudhar/go-book-store-user-api/logger"
)
var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info(" applications logs started ")
	router.Run(":8080")

}
