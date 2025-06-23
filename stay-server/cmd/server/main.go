package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	appPkg "stay-server/app"
)

func init() {
	log.Print("start init")
}

func main() {
	var app = appPkg.NewApp(1, gin.ReleaseMode)
	fmt.Print(app.InstanceId)

	app.GatewayInst.StartApiGateway()
}
