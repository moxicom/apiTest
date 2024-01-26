package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"testAPI/models"

	"github.com/gin-gonic/gin"
)

const (
	agePrefix    = "agify"
	genderPrefix = "genderize"
	nationPrefix = "nationalize"
)

func (h *handler) GetPeople(c *gin.Context) {
	nameFilter := c.Query("name")
	surnameFilter := c.Query("surname")
	patronymicFilter := c.Query("patronymic")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	person := models.Person{
		Name:       nameFilter,
		Surname:    surnameFilter,
		Patronymic: patronymicFilter,
	}

	filters := models.Filters{
		Limit:  limitStr,
		Offset: offsetStr,
	}

	people, err := h.service.People.GetPeople(person, filters)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, people)
}

func (h *handler) CreatePerson(c *gin.Context) {
	var input models.Person
	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(input)
	// insert new info
	info, err := getInfo(input.Name)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.Age = info.Age
	input.Country = info.Country
	input.Gender = info.Gender

	err = h.service.CreatePerson(&input)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, input)
}

func (h *handler) DeletePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.DeletePerson(id)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *handler) UpdatePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	var input models.PersonInput
	if err := c.BindJSON(&input); err != nil {
		if err != nil {
			newResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if err := h.service.UpdatePerson(id, input); err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
