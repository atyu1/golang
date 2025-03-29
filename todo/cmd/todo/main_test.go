package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = "todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building ...")
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build the tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running test ...")
	result := m.Run()

	fmt.Println("Cleaning up ...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "Test task 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTaskFromArgs", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)

		if err := cmd.Run(); err != nil {
			t.Error(err)
		}
	})

	task2 := "Test task 2"
	t.Run("AddNewTaskFromStdin", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdin, err := cmd.StdinPipe()
		if err != nil {
			t.Error(err)
		}
		io.WriteString(cmdStdin, task2)
		cmdStdin.Close()

		if err := cmd.Run(); err != nil {
			t.Error(err)
		}
	})

	t.Run("ListTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Error(err)
		}

		expected := fmt.Sprintf(" 1: %s\n 2: %s\n", task, task2)
		if expected != string(out) {
			t.Errorf("Expected: %q, got: %q", expected, string(out))
		}
	})

	t.Run("DeleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-delete", "1")

		if err := cmd.Run(); err != nil {
			t.Error(err)
		}

		cmd = exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Error(err)
		}

		expected := fmt.Sprintf(" 1: %s\n", task2)
		if expected != string(out) {
			t.Errorf("Expected: %q, got: %q", expected, string(out))
		}
	})
}
