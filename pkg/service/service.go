package service

import (
	"testAPI/models"
	"testAPI/pkg/repository"
)

type People interface {
	GetPeople(person models.Person, filters models.Filters) ([]models.Person, error)
	CreatePerson(person *models.Person) error
	DeletePerson(id int) error
	UpdatePerson(id int, person models.PersonInput) error
}

type Service struct {
	People
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		People: newPeopleService(repository.People),
	}
}
