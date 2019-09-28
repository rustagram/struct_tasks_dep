package models


import (
	"testing"
	_ "fmt"
	"github.com/jmoiron/sqlx"
)


var (
	cm ContactManagerI
	db *sqlx.DB
)

func TestContactManger(t *testing.T) {
	var err error
	cm, err = NewContactManager()
	if(err != nil){
		t.Error("Can't connect to database")
	}
}

func TestContactManagerAdd(t *testing.T){
	err := cm.Add(Contact{"Nurali", 23, "male", "2323"})

        if (err != nil) {
		t.Error("Contact not added")
	}
}


func TestContactMangerUpdate(t *testing.T) {

	err := cm.Update(3, Contact{"Nurali", 23, "male", "+998901221212"})

        if (err != nil) {
		t.Error("Contact not updated")
	}

}

func TestContactMangerGetAll(t *testing.T) {
	_, err := cm.GetAll()
        if (err != nil) {
		t.Error("Cannot get all contacts: ", err)
	}
}

func TestContactMangerDel(t *testing.T) {
	err := cm.Del(2)

        if (err != nil) {
		t.Error("Contact not deleted")
	}
}

func TestContactManagerListAll(t *testing.T) {
	err := cm.ListAll()
	if (err != nil) {
		t.Error("Helloo!")
	}
}
