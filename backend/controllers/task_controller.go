package controllers

import (
	"log"
	"net/http"

	"github.com/CoffeeCoder-Go/todolistapp/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateTaskController(ctx *gin.Context){
	var task models.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		log.Fatalf("Erro!, Ocorreu um erro ao validar JSON!Erro:%v.\n",err.Error())
		ctx.JSON(http.StatusBadRequest,gin.H{
			"err":err.Error(),
		})
		return
	} 

	task.GenerateID()

	DB.Create(&task)

	ctx.JSON(http.StatusCreated,gin.H{
		"msg":"Criado com sucesso!",
		"created":task,
	})

}

func GetAllTasksController(ctx *gin.Context){
	var tasks []models.Task

	DB.Find(&tasks)

	ctx.JSON(http.StatusOK,tasks)
}

func GetTaskByIDController(ctx *gin.Context){

}

func UpdateTaskController(ctx *gin.Context){

}

func DeleteATaskController(ctx *gin.Context){

} 