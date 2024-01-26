package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testAPI/models"
)

func makeRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getInfo(name string) (models.Person, error) {
	result := models.Person{}

	// get age
	ageData, err := makeRequest(fmt.Sprintf("https://api.%s.io/?name=%s", agePrefix, name))
	if err != nil {
		return models.Person{}, err
	}
	var ageResponse models.AgeResponse
	err = json.Unmarshal(ageData, &ageResponse)
	if err != nil {
		return models.Person{}, err
	}
	result.Age = ageResponse.Age

	// get gender
	genderData, err := makeRequest(fmt.Sprintf("https://api.%s.io/?name=%s", genderPrefix, name))
	if err != nil {
		return models.Person{}, err
	}
	var genderResponse models.GenderResponse
	err = json.Unmarshal(genderData, &genderResponse)
	if err != nil {
		return models.Person{}, err
	}
	result.Gender = genderResponse.Gender

	// get nation
	nationData, err := makeRequest(fmt.Sprintf("https://api.%s.io/?name=%s", nationPrefix, name))
	if err != nil {
		return models.Person{}, err
	}
	var nationResponse models.NationalityResponse
	err = json.Unmarshal(nationData, &nationResponse)
	if err != nil {
		return models.Person{}, err
	}
	result.Country = nationResponse.Country[0].CountryID // countries require at least 1 element

	return result, nil
}
