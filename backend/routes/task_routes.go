package routes

import (
	"github.com/CoffeeCoder-Go/todolistapp/backend/controllers"
	"github.com/gin-gonic/gin"
)

func NewTaskRouter(rg *gin.RouterGroup){
	tr := rg.Group("/tasks")

	tr.GET("/",controllers.GetAllTasksController)

	tr.POST("/",controllers.CreateTaskController)
}