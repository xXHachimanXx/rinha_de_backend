package schemas

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Person struct {
	ID        string         `gorm:"type:uuid;primaryKey"`
	Nickname  string         `gorm:"not null"`
	Name      string         `gorm:"not null"`
	Birthdate string         `gorm:"not null"`
	Stack     pq.StringArray `gorm:"type:text[]"`
}

func NewPerson() *Person {
	person := Person{
		ID: uuid.NewV4().String(),
	}

	return &person
}

func BuildPerson(nickname string, name string, birthdate string, stack []string) *Person {
	return &Person{
		ID:        uuid.NewV4().String(),
		Nickname:  nickname,
		Name:      name,
		Birthdate: birthdate,
		Stack:     pq.StringArray(stack),
	}
}
