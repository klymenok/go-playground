package mydb

import "log"

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
	createTaskQuery := `insert into task (title, description, created_by, assignee) values ('$1', '$2', $3, $4)`
	result, err := db.Exec(createTaskQuery, t.Title, t.Description, t.CreatedBy, t.Assignee)
	if err != nil {
		log.Fatalln(err)
	}
	t.Id, _ = result.LastInsertId()
}

func (t *Task) Update() {
	db := Connection()
	defer db.Close()

	updateTaskQuery := `update task set title='$1', description='$2', assignee=$3, completed=$4 where id=$5`
	_, err := db.Exec(updateTaskQuery, t.Title, t.Description, t.Assignee, t.Completed, t.Id)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetTaskById(taskId int64) (Task, error) {
	var task Task
	db := Connection()
	defer db.Close()

	getTaskQuery := `select * from task where id=$1`
	res := db.QueryRow(getTaskQuery, taskId)
	err := res.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedBy, &task.Assignee)
	if err != nil {
		return task, err
	}
	return task, nil
}

func DeleteTaskById(taskId int64) {
	db := Connection()
	defer db.Close()
	deleteTaskQuery := `delete from task where id=$1`
	db.Exec(deleteTaskQuery, taskId)
}
