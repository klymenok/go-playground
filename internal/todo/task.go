package todo

import (
	"database/sql"
	"fmt"
	"log"
)

type Task struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   int64  `json:"created_by"`
	Assignee    int64  `json:"assignee"`
	Completed   bool   `json:"completed"`
}

type tasks struct {
	db *sql.DB
}

func NewTask(db *sql.DB) *tasks {
	return &tasks{db}
}

func (t *tasks) Create(task *Task) {

	createTaskQuery := fmt.Sprintf(
		"insert into task (title, description, created_by, assignee) values ('%s', '%s', %d, %d)",
		task.Title,
		task.Description,
		task.CreatedBy,
		task.Assignee)
	result, err := t.db.Exec(createTaskQuery)
	if err != nil {
		log.Fatalln(err)
	}
	task.Id, _ = result.LastInsertId()
}

func (t *tasks) Update(task Task) {
	updateTaskQuery := fmt.Sprintf(
		"update task set title='%s', description='%s', assignee=%d, completed=%t where id=%d",
		task.Title,
		task.Description,
		task.Assignee,
		task.Completed,
		task.Id)
	_, err := t.db.Exec(updateTaskQuery)
	if err != nil {
		log.Fatalln(err)
	}
}

func (t *tasks) Complete(task *Task) {
	_, err := t.db.Exec("update task set completed=True where id=?", task.Id)
	if err != nil {
		log.Fatalln(err)
	}
	task.Completed = true
}

func (t *tasks) DeleteById(taskId int64) {
	deleteTaskQuery := fmt.Sprintf("delete from task where id=%d", taskId)
	t.db.Exec(deleteTaskQuery)
}

func (t *tasks) ById(taskId int64) (Task, error) {
	var task Task

	err := t.db.QueryRow(
		"select * from task where id=?", taskId,
	).Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.CreatedBy,
		&task.Assignee,
		&task.Completed,
	)
	if err != nil {
		return task, err
	}
	return task, nil
}
