package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ravayak/copee_backend/app/utils/logger"
	f "github.com/ravayak/copee_backend/datasources/firebase"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	router     = gin.Default()
	UsersByUID = make(map[string]mysql.User)
)

func init() {
	// Firebase initialization
	f.FirebaseInit()
	// Mysql Initialization
	mysql.MysqlInit("copee-v2")
}

func StartApp() {
	port := "8081"

	MapUrls()
	logger.Info("About to start the application on port : " + port)
	router.Run(":" + port)

}
