package todo_test

import (
	"testing"

	"github.com/atyu1/golang/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "Test task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "Test Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed, expected False")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("Complete should mark the task completed, its still not")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"Test task 1",
		"Test task 2",
		"Test task 3",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected after delete %q, got %q", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "Test task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l1[0].Task)
	}

	if err := l1.Save("testTaskL1.json"); err != nil {
		t.Fatalf("Cannot write to file: %s", err)
	}

	if err := l2.Get("testTaskL1.json"); err != nil {
		t.Fatalf("Cannot open file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("The loaded data does not match, expected %q, got %q", l1[0].Task, l2[0].Task)
	}
}
