package todo_test

import (
	"github.com/gabrielbussolo/todo"
	"os"
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
	t.Run("save and get file", func(t *testing.T) {
		l1 := todo.List{}
		l2 := todo.List{}

		taskName := "New Task"
		l1.Add(taskName)

		if l1[0].Task != taskName {
			t.Errorf("Expected task name %q, got %q", taskName, l1[0].Task)
		}

		temp, err := os.CreateTemp("", "")
		if err != nil {
			t.Fatalf("Error creating temp file: %s", err)
		}
		defer os.Remove(temp.Name())

		if err := l1.Save(temp.Name()); err != nil {
			t.Fatalf("Error saving list to a file: %s", err)
		}
		if err := l2.Get(temp.Name()); err != nil {
			t.Fatalf("Error getting list from temp file: %s", err)
		}
		if l1[0].Task != l2[0].Task {
			t.Errorf("Task %q should be equal to task %q", l1[0].Task, l2[0].Task)
		}
	})
}
