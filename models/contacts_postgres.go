package models

type ContactManager struct {
}

//Конструктор для структуры Менджер Контаков
func NewContactManager() (ContactManagerI, error) {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")

	if err != nil {
		return nil, err
	}

	cm := ContactManager{}

	return &cm, nil
}

//Функция добавления нового контакта в массив
func (c *ContactManager) Add(ct Contact) error {
	return nil
}

//Функция изменения определённого контакта
func (c *ContactManager) Update(i int, ct Contact) error {
	return nil
}

//Функция удаления контакта по айди
func (c *ContactManager) Del(i int) error {
	return nil
}

//Функция распечатки всех контактов
func (c *ContactManager) GetAll() ([]Contact, error) {
	return []Contact{}, nil
}

