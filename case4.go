package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Secret key to sign the token
var jwtSecret = []byte("mySecretKey")

// Struct for login credentials
type LoginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Struct for JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Function to generate JWT token
func generateToken(username string) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create claims with username and expiration time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Function to authenticate and issue JWT token
func login(c *gin.Context) {
	var credentials LoginCredentials

	// Bind JSON input to credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// For simplicity, use hardcoded username/password
	if credentials.Username != "admin" || credentials.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := generateToken(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return token as response
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Middleware to validate JWT
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// If token is missing
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Parse the JWT token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// Check if token is valid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Save the username in context for future handlers
		c.Set("username", claims.Username)

		// Proceed to the next handler
		c.Next()
	}
}

// Protected route accessible only with a valid JWT
func protected(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + username, "status": "Authorized"})
}

func main() {
	r := gin.Default()

	// Public route for login
	r.POST("/login", login)

	// Protected route with JWT middleware
	r.GET("/protected", authMiddleware(), protected)

	// Run the server
	r.Run(":8080")
}
