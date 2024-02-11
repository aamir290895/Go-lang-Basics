package labstack

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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
	// Create an Echo instance
	e := echo.New()

	// middleware for all apis
	e.Use(middleware)
	// Define your routes
	e.GET("/", hello)
	e.GET("/users/:id", getUser, validateToken)
	e.POST("/users", createUser, validateToken)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start the server
	e.Start(":8080")
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
func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := 123
		username := "john_doe"

		// Generate JWT token
		token, err := generateToken(userID, username)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to generate token")
		}

		log.Println(token)

		// Set the token in the response headers (or as needed)
		c.Response().Header().Set("Authorization", "Bearer "+token)

		// Continue processing by invoking the next handler
		return next(c)
	}
}

func validateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.String(http.StatusUnauthorized, "Authorization header missing")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			return c.String(http.StatusUnauthorized, "Invalid token")
		}

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			// Token is valid, you can access claims.UserID, claims.Username, etc.
			log.Printf("User ID: %d, Username: %s", claims.UserID, claims.Username)
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "Invalid token")
	}
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
