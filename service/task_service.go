package service

import (
	"context"
	"jeryhardianto/golang-tugas/model/web"
)

type TaskService interface {
	Create(ctx context.Context, request web.TaskCreateRequest) web.TaskResponeseAll
	Update(ctx context.Context, request web.TaskUpdateRequest) web.TaskResponeseAll
	DoneTask(ctx context.Context, request web.TaskDoneRequest) web.TaskResponeseAll
	FindById(ctx context.Context, taskId int) web.TaskResponeseAll
	FindAll(ctx context.Context) []web.TaskResponeseAll
}
