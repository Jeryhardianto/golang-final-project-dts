package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"io"
	"jeryhardianto/golang-tugas/app"
	"jeryhardianto/golang-tugas/controller"
	"jeryhardianto/golang-tugas/helper"
	"jeryhardianto/golang-tugas/repository"
	"jeryhardianto/golang-tugas/service"
	"log"
	"net/http"
	"os"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	//server := gin.Default()

	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository, db, validate)
	taskController := controller.NewTaskController(taskService)

	router := httprouter.New()

	router.GET("/api/tasks", taskController.FindAll)
	router.GET("/api/tasks/:taskId", taskController.FindById)
	router.POST("/api/tasks", taskController.Create)
	router.PUT("/api/tasks/:taskId", taskController.Update)
	router.PATCH("/api/tasks/:taskId", taskController.DoneTask)
	port := os.Getenv("PORT")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/", helloHandler)
	log.Println("Listing for" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

	server := http.Server{
		//Addr:    "localhost:3000",
		Addr:    "0.0.0.0",
		Handler: router,
	}

	err := server.ListenAndServe()

	helper.PanicIfError(err)

}
