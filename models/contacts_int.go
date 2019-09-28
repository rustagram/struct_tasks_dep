package models


//Структура "Контакт" хранит один контакт

type Contact struct{
	Name	string	`db:"name"`
	Age	int	`db:"age"`
	Gender	string	`db:"gender"`
	Number	string	`db:"number"`
}

//Структура "Менеджер Контактов" хранит массив структур "Контакт"
type ContactManagerI interface {
	Add (ct Contact) error
	Update (i int, ct Contact) error
	Del (i int) error
	GetAll () ([]Contact, error)
	ListAll () error
}
