package main

import (
	"LogoForCardsGin/config"
	"LogoForCardsGin/internal/db"
	"LogoForCardsGin/internal/repository"
	"LogoForCardsGin/internal/server"
	"LogoForCardsGin/internal/services"
	logging "LogoForCardsGin/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net"
)

func main() {
	logger := logging.GetLogger()
	route := gin.Default()
	config, err := config.GetConfig()
	if err != nil {
		log.Println(err)
		return
	}
	DB := db.GetDbConnection()

	newRepository := repository.NewRepository(DB, logger)
	newService := services.NewServices(newRepository, logger)
	newServer := server.NewHandler(route, newService, logger)

	newServer.Init()

	//route.GET("/", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{"message": "Connected"})
	//})
	//
	//route.GET("/get_users", controllers.GetUsers)
	//route.POST("/add_user", controllers.AddUser)
	//route.POST("/update_logo", controllers.UpdateLogoJson)
	//route.POST("/update_logo_multi", controllers.UpdateLogoMultipart)
	//route.POST("/change_status", controllers.ChangeStatus)

	address := net.JoinHostPort(config.LocalHost.Host, config.LocalHost.Port)

	err = route.Run(address)
	if err != nil {
		log.Println(err)
		return
	}
}
