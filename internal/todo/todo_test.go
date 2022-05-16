package todo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

func TestTaskByIdSuccess(t *testing.T) {
	// setup
	db, mock, _ := sqlmock.New()

	rows := sqlmock.NewRows(
		[]string{"id", "title", "description", "created_by", "assignee", "completed"},
	).AddRow(1, "Title", "Description", 1, 1, false)
	mock.ExpectQuery(".*").WillReturnRows(rows)

	// test
	todo := NewToDo(db)
	task, _ := todo.Tasks.ById(1)

	// asserts
	assert.NotNil(t, task)
	assert.Equal(t, int64(1), task.Id)
}
