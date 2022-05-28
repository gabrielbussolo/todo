package todo_test

import (
	"github.com/gabrielbussolo/todo"
	"testing"
)

func TestTodoList(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		l := todo.List{}

		taskName := "New Task"
		l.Add(taskName)

		if l[0].Task != taskName {
			t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
		}
	})
}
