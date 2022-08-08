package repository

import (
	"context"
	"database/sql"
	"errors"
	"jeryhardianto/golang-tugas/helper"
	"jeryhardianto/golang-tugas/model/domain"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

func (repository *TaskRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, task domain.Task) domain.Task {
	SQL := "INSERT INTO task(task, assignee, dateline, status) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, task.NameTask, task.Assignee, task.Dateline, 0)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	task.Id = int(id)
	return task
}

func (repository *TaskRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, task domain.Task) domain.Task {
	SQL := "UPDATE task SET task= ?, assignee= ?, dateline= ? WHERE id= ?"
	_, err := tx.ExecContext(ctx, SQL, task.NameTask, task.Assignee, task.Dateline, task.Id)
	helper.PanicIfError(err)

	return task
}

func (repository *TaskRepositoryImpl) DoneTask(ctx context.Context, tx *sql.Tx, task domain.Task) domain.Task {
	SQL := "UPDATE task SET status = ? WHERE id= ?"
	_, err := tx.ExecContext(ctx, SQL, 1, task.Id)
	helper.PanicIfError(err)

	return task
}

func (repository *TaskRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, taskId int) (domain.Task, error) {
	SQL := "SELECT id, task, assignee, dateline, status FROM task WHERE id= ?"
	rows, err := tx.QueryContext(ctx, SQL, taskId)

	helper.PanicIfError(err)
	defer rows.Close()
	//Slice
	task := domain.Task{}

	if rows.Next() {
		err := rows.Scan(&task.Id, &task.NameTask, &task.Assignee, &task.Dateline, &task.Status)
		helper.PanicIfError(err)
		return task, nil
	} else {
		return task, errors.New("Task is not found")
	}

}

func (repository *TaskRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Task {
	SQL := "SELECT id, task, assignee, dateline, status FROM task"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	//Slice
	var tasks []domain.Task

	for rows.Next() {
		task := domain.Task{}
		err := rows.Scan(&task.Id, &task.NameTask, &task.Assignee, &task.Dateline, &task.Status)
		helper.PanicIfError(err)

		tasks = append(tasks, task)
	}

	return tasks

}
