package models

import (
	"fmt"
	_ "database/sql"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type ContactManager struct {
	db *sqlx.DB
}


func NewContactManager() (ContactManagerI, error) {
	cm := ContactManager{}
	var err error
	cm.db, err = sqlx.Connect("postgres", "user=postgres password=123 dbname=contacts sslmode=disable")

	if err != nil {
		return nil, err
	}
	return &cm, nil
}

//Функция добавления нового контакта в массив
func (c *ContactManager) Add(ct Contact) error {
	insertNew := `INSERT INTO contact(name, age, gender, number) values($1, $2, $3, $4)`
	c.db.MustExec(insertNew, ct.Name, ct.Age, ct.Gender, ct.Number)
	return nil
}

//Функция изменения определённого контакта
func (c *ContactManager) Update(i int, ct Contact) error {
	updateContact := `UPDATE contact SET name = $1, age = $2, gender = $3, number = $4 where id = $5`
	c.db.MustExec(updateContact, ct.Name, ct.Age, ct.Gender, ct.Number, i)
	return nil
}

//Функция удаления контакта по айди
func (c *ContactManager) Del(i int) error {
	delContact := `DELETE FROM contact WHERE id = $1`
	c.db.MustExec(delContact, i)
	return nil
}

//Функция распечатки всех контактов
func (c *ContactManager) GetAll() ([]Contact, error) {
	var contacts []Contact
	err := c.db.Select(&contacts, "SELECT name, age, gender, number FROM contact")
	if(err != nil){
		return nil, err
	}
	return contacts, nil
}

func (c *ContactManager)ListAll() error{
	var contacts []Contact
	err := c.db.Select(&contacts, "SELECT name, age, gender, number FROM contact")
	if(err != nil){
		return err
	}
	fmt.Println(contacts)
	return nil
}
