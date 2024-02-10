package labstack

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Student struct {
	Name     string `json:"name"`
	RollNum  int    `json:"roll_num"`
	Class    string `json:"class"`
	Position string `json:"position"`
}

func Setup() {
	// Create an Echo instance
	e := echo.New()

	// Define your routes
	e.GET("/", hello)
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start the server
	e.Start(":8080")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	// Get the user ID from the URL parameter
	id := c.Param("id")
	return c.String(http.StatusOK, "User ID: "+id)
}

func createUser(c echo.Context) error {
	// Parse request body

	var st []Student
	query := struct {
		Name    string `json:"name"`
		RollNum int    `json:"roll_num"`
	}{}

	if err := c.Bind(&query); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request payload")
	}

	// Now you can use the 'query' struct to access the data from the request body

	st = append(st, Student{Name: "aamir", RollNum: 1, Class: "BE", Position: "1st"})
	return c.JSON(http.StatusCreated, &st)
}

func updateUser(c echo.Context) error {
	// Get the user ID from the URL parameter
	id := c.Param("id")
	// Parse request body
	name := c.FormValue("name")
	return c.String(http.StatusOK, "User updated - ID: "+id+", Name: "+name)
}

func deleteUser(c echo.Context) error {
	// Get the user ID from the URL parameter
	id := c.Param("id")
	return c.String(http.StatusOK, "User deleted - ID: "+id)
}
