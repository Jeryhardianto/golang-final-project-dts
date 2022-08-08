package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"jeryhardianto/golang-tugas/helper"
	"jeryhardianto/golang-tugas/model/domain"
	"jeryhardianto/golang-tugas/model/web"
	"jeryhardianto/golang-tugas/repository"
)

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTaskService(taskRepository repository.TaskRepository, DB *sql.DB, validate *validator.Validate) TaskService {
	return &TaskServiceImpl{
		TaskRepository: taskRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *TaskServiceImpl) Create(ctx context.Context, request web.TaskCreateRequest) web.TaskResponeseAll {
	//validasi request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	task := domain.Task{
		NameTask: request.NameTask,
		Assignee: request.Assignee,
		Dateline: request.Dateline,
		Status:   request.Status,
	}

	service.TaskRepository.Save(ctx, tx, task)

	return helper.ToTaskResponse(task)

}

func (service *TaskServiceImpl) Update(ctx context.Context, request web.TaskUpdateRequest) web.TaskResponeseAll {
	//validasi request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	task, err := service.TaskRepository.FindById(ctx, tx, request.Id)

	helper.PanicIfError(err)

	task.NameTask = request.NameTask
	task.Assignee = request.Assignee
	task.Dateline = request.Dateline

	service.TaskRepository.Update(ctx, tx, task)

	return helper.ToTaskResponse(task)

}

func (service *TaskServiceImpl) DoneTask(ctx context.Context, request web.TaskDoneRequest) web.TaskResponeseAll {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	task := domain.Task{
		Id: request.Id,
	}

	service.TaskRepository.DoneTask(ctx, tx, task)

	return helper.ToTaskResponse(task)
}

func (service *TaskServiceImpl) FindAll(ctx context.Context) []web.TaskResponeseAll {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tasks := service.TaskRepository.FindAll(ctx, tx)

	return helper.ToTaskResponses(tasks)

}
