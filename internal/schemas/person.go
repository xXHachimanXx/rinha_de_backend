package schemas

import uuid "github.com/satori/go.uuid"

type Person struct {
	ID        string   `valid:"uuidv4"`
	Nickname  string   `valid:"required" json:"apelido"`
	Name      string   `valid:"required" json:"nome"`
	Birthdate string   `valid:"required" json:"nascimento"`
	Stack     []string `valid:"required" json:"stack"`
}

func NewPerson() *Person {
	person := Person{
		ID: uuid.NewV4().String(),
	}

	return &person
}

func BuildPerson(id string, nickname string, name string, birthdate string, stack []string) *Person {
	return &Person{
		ID:        id,
		Nickname:  nickname,
		Name:      name,
		Birthdate: birthdate,
		Stack:     stack,
	}
}
