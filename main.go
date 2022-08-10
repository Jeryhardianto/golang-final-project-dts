package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"jeryhardianto/golang-tugas/app"
	"jeryhardianto/golang-tugas/controller"
	"jeryhardianto/golang-tugas/helper"
	"jeryhardianto/golang-tugas/repository"
	"jeryhardianto/golang-tugas/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository, db, validate)
	taskController := controller.NewTaskController(taskService)

	router := httprouter.New()

	router.GET("/api/tasks", taskController.FindAll)
	router.GET("/api/tasks/:taskId", taskController.FindById)
	router.POST("/api/tasks", taskController.Create)
	router.PUT("/api/tasks/:taskId", taskController.Update)
	router.PATCH("/api/tasks/:taskId", taskController.DoneTask)

	server := http.Server{
		//Addr:    "localhost:3000",
		Addr:    "https://golang-final-project.herokuapp.com",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
