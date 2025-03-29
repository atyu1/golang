package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

type Stringer interface {
	String() string
}

var verbose bool

func (l *List) Verbose(v bool) {
	verbose = v
}

var skipCompleted bool

func (l *List) SkipCompleted(s bool) {
	skipCompleted = s
}

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("the element %d is not in range of items. Does not exists", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("the element %d is not in range of items. Does not exists", i)
	}

	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := " "
		if t.Done {
			if skipCompleted {
				continue
			}
			prefix = "X "
		}
		if verbose {
			formatted += fmt.Sprintf("%s %s%d: %s\n", t.CreatedAt, prefix, k+1, t.Task)
		} else {
			formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
		}
	}

	return formatted
}
