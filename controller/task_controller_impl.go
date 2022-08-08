package controller

import (
	"github.com/julienschmidt/httprouter"
	"jeryhardianto/golang-tugas/helper"
	"jeryhardianto/golang-tugas/model/web"
	"jeryhardianto/golang-tugas/service"
	"net/http"
	"strconv"
)

type TaskControllerImpl struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) TaskController {
	return &TaskControllerImpl{
		TaskService: taskService,
	}
}

func (controller *TaskControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	taskCreateRequest := web.TaskCreateRequest{}
	helper.ReadFromRequestBody(request, &taskCreateRequest)

	taskResponse := controller.TaskService.Create(request.Context(), taskCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *TaskControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	taskUpdateRequest := web.TaskUpdateRequest{}
	helper.ReadFromRequestBody(request, &taskUpdateRequest)

	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanicIfError(err)

	taskUpdateRequest.Id = id

	controller.TaskService.Update(request.Context(), taskUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *TaskControllerImpl) DoneTask(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	taskDoneRequest := web.TaskDoneRequest{}
	helper.ReadFromRequestBody(request, &taskDoneRequest)

	taskId := params.ByName("taskId")
	id, err := strconv.Atoi(taskId)
	helper.PanicIfError(err)

	taskDoneRequest.Id = id

	controller.TaskService.DoneTask(request.Context(), taskDoneRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TaskControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	taskResponses := controller.TaskService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
