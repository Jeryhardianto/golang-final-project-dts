package web

type TaskCreateRequest struct {
	NameTask string `validate: "required"`
	Assignee string `validate: "required"`
	Dateline string `validate: "required"`
	Status   int
}
