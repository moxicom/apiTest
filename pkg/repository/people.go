package repository

import (
	"fmt"
	"strconv"
	"testAPI/models"

	"gorm.io/gorm"
)

type peopleRepository struct {
	db *gorm.DB
}

func newPeopleRepository(db *gorm.DB) *peopleRepository {
	return &peopleRepository{
		db: db,
	}
}

func (r *peopleRepository) CreatePerson(person *models.Person) error {
	return r.db.Create(&person).Error
}

func (r *peopleRepository) DeletePerson(id int) error {
	return r.db.Delete(&models.Person{}, id).Error
}

func (r *peopleRepository) UpdatePerson(id int, person models.PersonInput) error {
	var temp models.Person
	if err := r.db.First(&temp, id).Error; err != nil {
		return err
	}

	temp.Name = person.Name
	temp.Surname = person.Surname
	if person.Patronymic != "" {
		temp.Patronymic = person.Patronymic
	}

	return r.db.Save(&temp).Error
}

func (r *peopleRepository) GetPeople(person models.Person, filters models.Filters) ([]models.Person, error) {
	var people []models.Person

	query := r.db.Model(&models.Person{})

	// Apply filters
	if person.Name != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%s%%", person.Name))
	}
	if person.Surname != "" {
		query = query.Where("surname LIKE ?", fmt.Sprintf("%s%%", person.Surname))
	}
	if person.Patronymic != "" {
		query = query.Where("patronymic LIKE ?", fmt.Sprintf("%s%%", person.Patronymic))
	}

	// Apply limit and offset
	if filters.Limit != "" {
		// Parse limit string to integer
		// Handle error accordingly
		limit, err := strconv.Atoi(filters.Limit)
		if err == nil {
			query = query.Limit(limit)
		}
	}
	if filters.Offset != "" {
		// Parse offset string to integer
		// Handle error accordingly
		offset, err := strconv.Atoi(filters.Offset)
		if err == nil {
			query = query.Offset(offset)
		}
	}

	if err := query.Find(&people).Error; err != nil {
		return nil, err
	}

	return people, nil
}
