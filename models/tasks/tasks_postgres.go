package tasks

import(
	_ "database/sql"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type TaskManager struct{
	db *DB.sqlx
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
	addNewTask = `INSERT INTO tasks VALUES($1, $2, $3, $4, $5)`
	tm.db.MustExec(addNewTask, t.Assignee, t.Title, t.Deadline, t.Done, t.Overdue)
	return nil
}

func (tm *TaskManager) Assign (name string, id int) error {
	assign = `UPDATE tasks SET assignee = $1 where id = $2`
	tm.db.MustExec(assign, name, id)
	return nil
}

func (tm *TaskManager) MakeLate () error {
	makeLate = `UPDATE tasks SET overdue=true where done=false and deadline<current_date`
	tm.db.MustExec(makeLate)
	return nil
}

