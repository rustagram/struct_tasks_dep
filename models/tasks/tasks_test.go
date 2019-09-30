package models

import (
	"testing"
	"time"
	"github.com/jmoiron/sqlx"
)

var (
	tm TaskManagerI
	db *sqlx.DB
)

func TestNewTaskManager (t *testing.T) {
	var err error
	tm, err = NewTaskManager()
	if err != nil{
		t.Error("Can't connect to database")
	}
}

func TestTaskManagerAdd (t *testing.T) {
	tt, _ := time.Parse("2006-02-01", "2019-10-01")
	err := tm.Add(Task{"Nurali", "Contacts", tt, false, false})
	if err != nil{
		t.Error("Can't add a contact")
	}
}

func TestTaskManagerAssign (t *testing.T) {
	err := tm.Assign("Nurali", 2)
	if err != nil{
		t.Error("Can't assign a task")
	}
}

func TestTaskManagerMakeLate (t *testing.T) {
	err := tm.MakeLate()
	if err != nil{
		t.Error("Can't make late tasks")
	}
}

func TestTaskManagerMakeDone (t *testing.T) {
	err := tm.MakeDone(2)
	if err != nil{
		t.Error("Can't make done a task")
	}
}

func TestTaskManagerUpdate (t *testing.T) {
	tt, _ := time.Parse("2006-02-01", "2019-07-01")
	err := tm.Update(Task{"Timur", "Lightom", tt, false, true}, 1)
	if err != nil{
		t.Error("Can't update a task")
	}
}

func TestTaskManagerDel (t *testing.T) {
	err := tm.Del(3)
	if err != nil{
		t.Error("Can't delete a task")
	}
}

func TestTaskManagerGetAll (t *testing.T) {
	_, err := tm.GetAll()
	if err != nil{
		t.Error("Can't get all contacts")
	}
}
