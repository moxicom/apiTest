package models

// nameFilter := c.Query("name")
// surnameFilter := c.Query("surname")
// patronymic := c.Query("patronymic")
// limitStr := c.Query("limit")
// offsetStr := c.Query("offset")

type Filters struct {
	Limit  string
	Offset string
}
