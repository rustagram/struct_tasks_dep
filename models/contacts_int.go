package models

//Структура "Контакт" хранит один контакт

type Contact struct{
	name string
	age int
	gender string
	number string
}

//Структура "Менеджер Контактов" хранит массив структур "Контакт"
type ContactManagerI interface {
	Add (ct Contact) error
	Update (i int, ct Contact) error
	Del (i int) error
	GetAll () ([]Contact, error)
}


