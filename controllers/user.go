package controllers

import (
	"net/http"
	"time"

	"financial-journey/database"
	"financial-journey/helper"
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
		helper.RespondWithError(c, http.StatusBadRequest, "password harus 8 karakter", nil)
		return
	}
	if len(user.Password) > 8 {
		helper.RespondWithError(c, http.StatusBadRequest, "password harus 8 karakter", nil)
		return
	}

	//cek user role apakah admin atau user
	if user.Role != "admin" && user.Role != "user" {
		helper.RespondWithError(c, http.StatusBadRequest, "role tidak ditemukan", nil)
		return
	}
	//userList := repository.GetUserByName(database.DbConnection, user.Username)
	userExist, err := repository.GetUserByUsername(database.DbConnection,user.Username);

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userExist != nil {
		helper.RespondWithError(c, http.StatusBadRequest, "Username already exists", nil)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, "gagal melakukan hashing password", err)
		return
	}
	user.Password = string(hashedPassword)

	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}
	helper.RespondWithSuccess(c, http.StatusOK, "success register user", nil)


}
func LoginHandler(c *gin.Context) {
		var loginData structs.LoginData
		if err := c.ShouldBindJSON(&loginData); err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, "kesalahan payload body", err)
			return
		}

		user, err := repository.GetUserByUsername(database.DbConnection,loginData.Username)

		if err != nil {
			helper.RespondWithError(c, http.StatusInternalServerError, "kesalahan sistem", err)
			return
		}

		if user == nil {
			helper.RespondWithError(c, http.StatusUnauthorized, "Invalid username", nil)
			return
		}

		if user.Username != loginData.Username {
			helper.RespondWithError(c, http.StatusUnauthorized, "Invalid username", nil)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, "Invalid password", err)
			return
		}

		token, err := generateToken(user)
		if err != nil {
			helper.RespondWithError(c, http.StatusInternalServerError, "Failed to generate token", err)
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, "success login", token)
		
	
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

	helper.RespondWithSuccess(c, http.StatusOK, "success get user", gin.H{"username": username, "role": role, "id": id})
}

// adminHandler handles requests from admin users
func AdminHandler(c *gin.Context) {
	username := c.GetString("username")
	role := c.GetString("role")
	id := c.GetUint("id")

	if role != "admin" {
		helper.RespondWithError(c, http.StatusInternalServerError, "Permission denied", nil)
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, "success get admin", gin.H{"username": username, "role": role, "id": id})
}