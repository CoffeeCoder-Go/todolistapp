package main

import (
	"github.com/CoffeeCoder-Go/todolistapp/backend/config"
	"github.com/CoffeeCoder-Go/todolistapp/backend/controllers"
	"github.com/CoffeeCoder-Go/todolistapp/backend/models"
	"github.com/CoffeeCoder-Go/todolistapp/backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var rt *gin.Engine = gin.Default()
var db *gorm.DB

func defineInitials(){
	rt.Use(cors.Default())
	db = config.ConnectToDB("./db/dev.db")
	config.MigrateAllTables(db,&models.Task{})
	controllers.DB = db
}

func main(){
	routes.NewTaskRouter(&rt.RouterGroup)
	defineInitials()
	rt.Run(":3333")
}