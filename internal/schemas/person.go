package schemas

import uuid "github.com/satori/go.uuid"

type Person struct {
	ID        string `gorm:"primaryKey"`
	Nickname  string
	Name      string
	Birthdate string
	Stack     string `gorm:"type:text"`
}

func NewPerson() *Person {
	person := Person{
		ID: uuid.NewV4().String(),
	}

	return &person
}

func BuildPerson(nickname string, name string, birthdate string, stack string) *Person {
	return &Person{
		ID:        uuid.NewV4().String(),
		Nickname:  nickname,
		Name:      name,
		Birthdate: birthdate,
		Stack:     stack,
	}
}
