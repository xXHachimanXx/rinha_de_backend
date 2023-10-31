package schemas

import (
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Person struct {
	ID        string         `gorm:"type:uuid;primaryKey"`
	Nickname  string         `gorm:"not null"`
	Name      string         `gorm:"not null"`
	Birthdate time.Time      `gorm:"not null"`
	Stack     pq.StringArray `gorm:"type:text[]"`
}

func (Person) TableName() string {
	return "person"
}

func NewPerson() *Person {
	person := Person{
		ID: uuid.NewV4().String(),
	}

	return &person
}

func BuildPersonRequest(nickname string, name string, birthdate time.Time, stack []string) *Person {

	return &Person{
		ID:        uuid.NewV4().String(),
		Nickname:  nickname,
		Name:      name,
		Birthdate: birthdate,
		Stack:     pq.StringArray(stack),
	}
}
