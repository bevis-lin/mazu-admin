package main

import (
	"mazu/admin/controllers"
	"mazu/admin/middleware"
	"mazu/admin/service"
	"net/http"
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

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controllers.LoginController = controllers.LoginHandler(loginService, jwtService)

	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	r.GET("/mint/requests", controllers.GetAllRequests)

	r.Use(middleware.AuthorizeJWT())

	r.POST("/mint/requests/approve", controllers.ApproveMintRequest)

	r.Run(":8081")

}
