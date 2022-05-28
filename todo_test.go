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
	t.Run("complete", func(t *testing.T) {
		l := todo.List{}

		taskName := "New Task"
		l.Add(taskName)

		if l[0].Task != taskName {
			t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
		}
		if l[0].Done {
			t.Errorf("New task should not be completed")
		}
		l.Complete(1)
		if !l[0].Done {
			t.Errorf("New task should be completed")
		}
	})
	t.Run("delete", func(t *testing.T) {
		l := todo.List{}

		tasks := []string{
			"Task 1",
			"Task 2",
			"Task 3",
		}

		for _, task := range tasks {
			l.Add(task)
		}

		if l[0].Task != tasks[0] {
			t.Errorf("Expected %q, got %q", tasks[0], l[0].Task)
		}

		l.Delete(2)

		if len(l) != 2 {
			t.Errorf("Expected list length %d, got %d", 2, len(l))
		}

		if l[1].Task != tasks[2] {
			t.Errorf("Expected %q, got %q", tasks[2], l[1].Task)
		}
	})
}
