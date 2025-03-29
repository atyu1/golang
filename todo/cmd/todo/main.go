package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/atyu1/golang/todo"
)

var todoFileName = "todo.json"

func main() {
	// Parsing command line arguments
	add := flag.Bool("add", false, "Add task to be included in todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Complete specific task")
	delete := flag.Int("delete", 0, "Delete specific task")
	verbose := flag.Bool("verbose", false, "Add date/time option")
	skipCompleted := flag.Bool("skip", false, "Don't show completed items in output")

	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Println("Creating todo file: ", todoFileName)
		_, err := os.Create(todoFileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	l.Verbose(*verbose) // Set verbose for outputs
	l.SkipCompleted(*skipCompleted)

	switch {
	case *list:
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(3)
		}

		l.Add(t)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(4)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(5)
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	var tasks []string{}

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	for s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return s.Text(), nil
}
