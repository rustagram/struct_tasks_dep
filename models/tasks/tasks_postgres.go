package models

import(
	_ "database/sql"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type TaskManager struct{
	db *sqlx.DB
}

func NewTaskManager() (TaskManagerI, error) {
	tm := TaskManager{}

	var err error
	tm.db, err = sqlx.Connect("postgres", "user=postgres password=123 dbname=tasks sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &tm, nil
}

func (tm *TaskManager) Add(t Task) error {
	addNewTask := `INSERT INTO tasks(assignee, title, deadline, done, overdue) VALUES($1, $2, $3, $4, $5)`
	tm.db.MustExec(addNewTask, t.Assignee, t.Title, t.Deadline, t.Done, t.Overdue)
	return nil
}

func (tm *TaskManager) Assign (name string, id int) error {
	assign := `UPDATE tasks SET assignee = $1 where id = $2`
	tm.db.MustExec(assign, name, id)
	return nil
}

func (tm *TaskManager) MakeLate () error {
	makeLate := `UPDATE tasks SET overdue=true where done=false and deadline<current_date`
	tm.db.MustExec(makeLate)
	return nil
}

func (tm *TaskManager) MakeDone (id int) error {
	makeDone := `UPDATE tasks SET done=true WHERE id=$1`
	tm.db.MustExec(makeDone, id)
	return nil
}

func (tm *TaskManager) Update (t Task, id int) error {
	update := `UPDATE tasks SET assignee=$1, title=$2, deadline=$3, done=$4, overdue=$5 where id=$6`
	tm.db.MustExec(update, t.Assignee, t.Title, t.Deadline, t.Done, t.Overdue, id)
	return nil
}

func (tm *TaskManager) Del (id int) error {
	del := `DELETE FROM tasks WHERE id=$1`
	tm.db.MustExec(del, id)
	return nil
}

func (tm * TaskManager) GetAll () ([]Task, error) {
	var tasks []Task
	err := tm.db.Select(&tasks, "SELECT assignee, title, deadline, done, overdue FROM tasks")
	if err != nil{
		return nil, err
	}
	return tasks, nil
}
