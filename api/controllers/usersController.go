package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	initializaers "github.com/lucapierini/api/initializers"
	"github.com/lucapierini/api/models"
	"golang.org/x/crypto/bcrypt"
	// "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"os"
)

func Signup(c *gin.Context){
	// Obtener el email y la contraseña del cuerpo de la request
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// Hashear la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	// Crear el usuario
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializaers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	// Responder con el usuario creado
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context){
	// Obtener el email y la contraseña del cuerpo de la request
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// Buscar el usuario por email
	var user models.User
	initializaers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	// Comparar la contraseña hasheada con la contraseña del cuerpo de la request
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	// Generar el jwt token
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Firmar el token con una clave secreta
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create token"})
		return
	}

	// Responder con el token
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30 ,"", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message":"Welcome "+body.Email})
}

func Validate(c *gin.Context){
	user, _ := c.Get("user")


	c.JSON(http.StatusOK, gin.H{"message": user})
}