package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Id      int32
	Dao     *gorm.DB
	Gateway *gin.Engine
}

func NewApp() *App {

	return &App{
		Id: 1,
	}
}
