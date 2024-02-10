package ginframework

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name     string `json:"name"`
	RollNum  int    `json:"roll_num"`
	Class    string `json:"class"`
	Position string `json:"position"`
}

func main() {
	// Create a Gin router
	r := gin.Default()

	// Define your routes
	r.GET("/", hello)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	// Start the server
	r.Run(":8080")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func getUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	id := c.Param("id")
	c.String(http.StatusOK, "User ID: "+id)
}

func createUser(c *gin.Context) {
	// Parse request body
	var st []Student
	var query struct {
		Name    string `json:"name"`
		RollNum int    `json:"roll_num"`
	}

	if err := c.BindJSON(&query); err != nil {
		c.String(http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Now you can use the 'query' struct to access the data from the request body
	st = append(st, Student{Name: query.Name, RollNum: query.RollNum, Class: "BE", Position: "1st"})
	c.JSON(http.StatusCreated, &st)
}

func updateUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	id := c.Param("id")
	// Parse request body
	name := c.PostForm("name")
	c.String(http.StatusOK, "User updated - ID: "+id+", Name: "+name)
}

func deleteUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	id := c.Param("id")
	c.String(http.StatusOK, "User deleted - ID: "+id)
}
