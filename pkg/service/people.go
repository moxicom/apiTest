package service

import (
	"testAPI/models"
	"testAPI/pkg/repository"
)

type peopleService struct {
	repository repository.People
}

func newPeopleService(repository repository.People) *peopleService {
	return &peopleService{
		repository: repository,
	}
}

func (s *peopleService) GetPeople(person models.Person, filters models.Filters) ([]models.Person, error) {
	return s.repository.GetPeople(person, filters)
}
func (s *peopleService) CreatePerson(person *models.Person) error {
	return s.repository.CreatePerson(person)
}

func (s *peopleService) DeletePerson(id int) error {
	return s.repository.DeletePerson(id)
}

func (s *peopleService) UpdatePerson(id int, person models.PersonInput) error {
	return s.repository.UpdatePerson(id, person)
}
