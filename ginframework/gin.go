package ginframework

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Student struct {
	Name     string `json:"name"`
	RollNum  int    `json:"roll_num"`
	Class    string `json:"class"`
	Position string `json:"position"`
}

type CustomClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte("your-secret-key")

func Setup() {
	// Create a Gin router
	r := gin.Default()

	// Define your routes
	r.GET("/", hello)
	r.GET("/users/:id", getUser, validateToken())
	r.POST("/users", createUser, validateToken())
	r.PUT("/users/:id", updateUser, validateToken())
	r.DELETE("/users/:id", deleteUser)

	// Start the server
	r.Run(":8080")
}

func generateToken(userID int, username string) (string, error) {
	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
		},
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.String(http.StatusUnauthorized, "Authorization header missing")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.String(http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			// Token is valid, you can access claims.UserID, claims.Username, etc.
			log.Printf("User ID: %d, Username: %s", claims.UserID, claims.Username)
			c.Next()
		} else {
			c.String(http.StatusUnauthorized, "Invalid token")
			c.Abort()
		}
	}
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func getUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	var err error
	id := c.Param("id")

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("error occoured"))
	}
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
