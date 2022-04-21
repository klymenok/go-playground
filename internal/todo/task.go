package todo

import (
	"fmt"
	"log"

	db "github.com/klymenok/go-playground/internal/db"
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

func NewTask(db *db.DB) *Task {
	task := &Task{}
	task.db = db
	return task
}

func (t *Task) Create() {
	// db := Connection()
	// defer db.Close()
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
	// db := Connection()
	// defer db.Close()

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
	// db := Connection()
	// defer db.Close()
	updateTaskQuery := fmt.Sprintf(
		"update task set completed=True where id=%d",
		t.Id)
	_, err := t.db.Exec(updateTaskQuery)
	if err != nil {
		log.Fatalln(err)
	}
	t.Completed = true
}

func GetTaskById(taskId int64) (Task, error) {
	var task Task
	// db := Connection()
	// defer db.Close()

	// getTaskQuery := fmt.Sprintf("select * from task where id=%d", taskId)
	// res := t.db.QueryRow(getTaskQuery)
	// err := res.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedBy, &task.Assignee, &task.Completed)
	// if err != nil {
	//   return task, err
	// }
	return task, nil
}

func DeleteTaskById(taskId int64) {
	// db := Connection()
	// defer db.Close()
	// deleteTaskQuery := fmt.Sprintf("delete from task where id=%s", taskId)
	// db.Exec(deleteTaskQuery)
}
