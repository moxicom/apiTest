# apiTest
Used clean architecture.

Database: postgres

REST framework: gin

ORM: gorm

To make migration use `make migration` command.
To run server use `make run` command

Get request parameters: name, surname, patronymic, offset, limit

Post/put json data:
```
type Person struct {
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic"`
}
```
