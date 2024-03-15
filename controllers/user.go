package controllers

import (
	"net/http"
	"time"

	"financial-journey/database"
	"financial-journey/repository"
	"financial-journey/structs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	users     []structs.User
	secretKey = []byte("secret")
)

func RegisterHandler(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		
		panic(err)
	}
	//cek user password lenght
	if len(user.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password harus 8 karakter"})
		return
	}
	if len(user.Password) > 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password harus 8 karakter"})
		return
	}

	//cek user role apakah admin atau user
	if user.Role != "admin" && user.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role tidak ditemukan"})
		return
	}
	//userList := repository.GetUserByName(database.DbConnection, user.Username)
	userExist, err := repository.GetUserByUsername(database.DbConnection,user.Username);

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userExist != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal melakukan hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "success register user",
	})


}
func LoginHandler(c *gin.Context) {
		var loginData structs.LoginData
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := repository.GetUserByUsername(database.DbConnection,loginData.Username)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
			return
		}

		if user.Username != loginData.Username {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		token, err := generateToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message":"sukses login","token": token})
	
}

func generateToken(user *structs.User) (string, error) {
	claims := &structs.JWTClaims{
		Username: user.Username,
		Role:     user.Role,
		ID:       user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserHandler(c *gin.Context) {
	username := c.GetString("username")
	role := c.GetString("role")
	id := c.GetUint("id")
	
	c.JSON(http.StatusOK, gin.H{"username": username, "role": role, "id": id})
}

// adminHandler handles requests from admin users
func AdminHandler(c *gin.Context) {
	role := c.GetString("role")

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome, admin!"})
}