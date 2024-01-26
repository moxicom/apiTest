package repository

import (
	"testAPI/models"

	"gorm.io/gorm"
)

type People interface {
	CreatePerson(person *models.Person) error
	DeletePerson(id int) error
	UpdatePerson(id int, person models.PersonInput) error
	GetPeople(person models.Person, filters models.Filters) ([]models.Person, error)
}

type Repository struct {
	People
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		People: newPeopleRepository(db),
	}
}
