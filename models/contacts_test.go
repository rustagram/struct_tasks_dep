package models


import (
	"testing"
)


var (
	cm ContactManagerI
)

func TestContactMangerAdd(t *testing.T) {
	cm = NewContactManager()

	err := cm.Add(Contact{"Rustam", 23, "male", "+998908082596"})

        if (err != nil) {
		t.Error("Contact not added")
	}
}


func TestContactMangerUpdate(t *testing.T) {

	err := cm.Update(0, Contact{"Nurali", 23, "male", "+998901221212"})

        if (err != nil) {
		t.Error("Contact not updated")
	}

}

func TestContactMangerGetAll(t *testing.T) {
	_, err := cm.GetAll()
        if (err != nil) {
		t.Error("Cannot get all contacts")
	}
}

func TestContactMangerDel(t *testing.T) {
	err := cm.Del(0)

        if (err != nil) {
		t.Error("Contact not deleted")
	}
}

