package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"stay-server/internal/services/user"
)

func (this *GatewayApp) StartApiGateway() {
	var userServices user.UserServices

	commonPrefixGroup := this.Router.Group("/api")

	apiV1 := commonPrefixGroup.Group("/v1")

	apiV1Public := apiV1.Group("/public")
	{
		apiV1Public.GET("/login", userServices.Login)
		apiV1Public.POST("/register")
	}

	apiV1ProtectedV0 := apiV1.Group("/user")
	{
		apiV1ProtectedV0.GET("/data", func(context *gin.Context) {
			context.JSON(http.StatusOK, "user")
		})
	}

	apiV1ProtectedV1 := apiV1.Group("/master")
	{
		apiV1ProtectedV1.GET("/data", func(context *gin.Context) {
			context.JSON(http.StatusOK, "master")
		})
	}

	apiV1ProtectedV10 := apiV1.Group("/admin")
	{
		apiV1ProtectedV10.GET("/data", func(context *gin.Context) {
			context.JSON(http.StatusOK, "admin")
		})
	}

	publicApi := this.Router.Group("/api/public")
	{
		publicApi.POST("/user/login")
		publicApi.POST("/user/register")
	}

	log.Println("start")
	if err := this.Router.Run(":8088"); err != nil {
		log.Println(err)
	}

}
