package todo

import (
	"fmt"
	"log"

	"github.com/klymenok/go-playground/internal/db"
)

type Task struct {
	db *db.DB

	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   int64  `json:"created_by"`
	Assignee    int64  `json:"assignee"`
	Completed   bool   `json:"completed"`
}

func NewTask() *Task {
	task := &Task{}
	task.db = db.New()
	return task
}

func (t *Task) Create() {
	createTaskQuery := fmt.Sprintf(
		"insert into task (title, description, created_by, assignee) values ('%s', '%s', %d, %d)",
		t.Title,
		t.Description,
		t.CreatedBy,
		t.Assignee)
	result, err := t.db.Exec(createTaskQuery)
	if err != nil {
		log.Fatalln(err)
	}
	t.Id, _ = result.LastInsertId()
}

func (t *Task) Update() {
	updateTaskQuery := fmt.Sprintf(
		"update task set title='%s', description='%s', assignee=%d, completed=%t where id=%d",
		t.Title,
		t.Description,
		t.Assignee,
		t.Completed,
		t.Id)
	_, err := t.db.Exec(updateTaskQuery)
	if err != nil {
		log.Fatalln(err)
	}
}

func (t *Task) Complete() {
	updateTaskQuery := fmt.Sprintf(
		"update task set completed=True where id=%d",
		t.Id)
	_, err := t.db.Exec(updateTaskQuery)
	if err != nil {
		log.Fatalln(err)
	}
	t.Completed = true
}
