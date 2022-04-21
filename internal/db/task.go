package mydb

import (
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

func (t *Task) Create() {
	db := Connection()
	defer db.Close()
	createTaskQuery := fmt.Sprintf(
		"insert into task (title, description, created_by, assignee) values ('%s', '%s', %d, %d)",
		t.Title,
		t.Description,
		t.CreatedBy,
		t.Assignee)
	result, err := db.Exec(createTaskQuery)
	if err != nil {
		log.Fatalln(err)
	}
	t.Id, _ = result.LastInsertId()
}

func (t *Task) Update() {
	db := Connection()
	defer db.Close()

	updateTaskQuery := fmt.Sprintf(
		"update task set title='%s', description='%s', assignee=%d, completed=%t where id=%d",
		t.Title,
		t.Description,
		t.Assignee,
		t.Completed,
		t.Id)
	_, err := db.Exec(updateTaskQuery)
	if err != nil {
		log.Fatalln(err)
	}
}

func (t *Task) Complete() {
	db := Connection()
	defer db.Close()
	updateTaskQuery := fmt.Sprintf(
		"update task set completed=True where id=%d",
		t.Id)
	_, err := db.Exec(updateTaskQuery)
	if err != nil {
		log.Fatalln(err)
	}
	t.Completed = true
}

func GetTaskById(taskId int64) (Task, error) {
	var task Task
	db := Connection()
	defer db.Close()

	getTaskQuery := fmt.Sprintf("select * from task where id=%d", taskId)
	res := db.QueryRow(getTaskQuery)
	err := res.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedBy, &task.Assignee, &task.Completed)
	if err != nil {
		return task, err
	}
	return task, nil
}

func DeleteTaskById(taskId int64) {
	db := Connection()
	defer db.Close()
	deleteTaskQuery := fmt.Sprintf("delete from task where id=%s", taskId)
	db.Exec(deleteTaskQuery)
}
