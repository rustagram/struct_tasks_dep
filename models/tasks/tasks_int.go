package tasks

import(
	"time"
)

type struct Task {
	Assignee string
	Title string
	Deadline time.Time
	Done bool
	Overdue bool
}

type TaskManagerI interface{
	Add (tm Task) error
	Assign (name string, id int) error
	MakeLate () error
	MakeDone (id int) error
	Update (tm Task, id int) error
	Del (id int) error
	GetAll() ([]Task, error)
}
