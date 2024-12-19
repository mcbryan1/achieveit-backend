package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcbryan1/achieveit-backend/helpers"
	"github.com/mcbryan1/achieveit-backend/initializers"
	"github.com/mcbryan1/achieveit-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	_, user, tokenString, err := helpers.ProcessLogin(c)
	if err != nil {
		return
	}

	userResponse := helpers.CreateUserResponse(user)
	helpers.RespondWithSuccess(c, http.StatusOK, "Login successful", "000", gin.H{
		"token": tokenString,
		"user":  userResponse,
	})
}

func RegisterUser(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request", "001")
		return
	}

	if err := helpers.ValidateRequest(req, "User"); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "001")
		return
	}

	username := req["username"].(string)
	if helpers.UserExists(username) {
		helpers.RespondWithError(c, http.StatusBadRequest, "User already exists", "001")
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Error hashing password", "001")
		return
	}

	newUser := models.User{
		Username: req["username"].(string),
		Password: string(hashPassword),
	}

	if err := initializers.DB.Create(&newUser).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Error creating user", "001")
		return
	}

	helpers.RespondWithSuccess(c, http.StatusOK, "User created successfully", "000")
}
