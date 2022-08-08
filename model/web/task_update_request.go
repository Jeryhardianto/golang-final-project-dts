package web

type TaskUpdateRequest struct {
	Id       int    `validate: "required"`
	NameTask string `validate: "required"`
	Assignee string `validate: "required"`
	Dateline string `validate: "required"`
}
