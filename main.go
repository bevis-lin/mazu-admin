package main

import (
	"mazu/admin/controllers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	controllers.AdminAddress = os.Getenv("ADMIN_ADDRESS")
	controllers.AdminPrivateKey = os.Getenv("ADMIN_PRIVATE_KEY")
	controllers.Node = os.Getenv("ACCESS_NODE")

	r := gin.Default()

	r.GET("/mint/requests", controllers.GetAllRequests)

	r.POST("/mint/requests/approve", controllers.ApproveMintRequest)

	r.Run(":8081")

}
