package helper

import (
	"jeryhardianto/golang-tugas/model/domain"
	"jeryhardianto/golang-tugas/model/web"
)

func ToTaskResponse(task domain.Task) web.TaskResponeseAll {
	return web.TaskResponeseAll{
		Id:       task.Id,
		NameTask: task.NameTask,
		Assignee: task.Assignee,
		Dateline: task.Dateline,
		Status:   task.Status,
	}
}

func ToTaskResponses(tasks []domain.Task) []web.TaskResponeseAll {
	var taskResponse []web.TaskResponeseAll
	for _, task := range tasks {
		taskResponse = append(taskResponse, ToTaskResponse(task))
	}

	return taskResponse
}
