package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mcbryan1/achieveit-backend/helpers"
	"github.com/mcbryan1/achieveit-backend/initializers"
	"github.com/mcbryan1/achieveit-backend/models"
)

func CreateGoal(c *gin.Context) {
	userID, ok, err := helpers.GetUserIDFromContext(c)
	if !ok || err != nil {
		if err == nil {
			helpers.RespondWithError(c, http.StatusUnauthorized, "User not authenticated", "001")
		} else {
			helpers.RespondWithError(c, http.StatusInternalServerError, err.Error(), "500")
		}
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request", "001")
		return
	}

	if err := helpers.ValidateRequest(req, "Goal"); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "001")
		return
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Invalid user ID", "500")
		return
	}

	newGoal := models.Goal{
		UserID:      uid,
		Title:       req["title"].(string),
		Description: req["description"].(string),
		Progress:    0,
	}

	results := initializers.DB.Create(&newGoal)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	goalResponse := helpers.CreateGoalResponse(newGoal)
	helpers.RespondWithSuccess(c, http.StatusOK, "Goal created successfully", "000", goalResponse)
}

func GetGoals(c *gin.Context) {
	userID, ok, err := helpers.GetUserIDFromContext(c)
	if !ok || err != nil {
		if err == nil {
			helpers.RespondWithError(c, http.StatusUnauthorized, "User not authenticated", "001")
		} else {
			helpers.RespondWithError(c, http.StatusInternalServerError, err.Error(), "500")
		}
		return
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Invalid user ID", "500")
		return
	}

	var goals []models.Goal
	results := initializers.DB.Where("user_id = ?", uid).Preload("Milestones").Find(&goals)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	goalsResponse := helpers.FetchGoalsResponse(goals)
	helpers.RespondWithSuccess(c, http.StatusOK, "Goals fetched successfully", "000", goalsResponse)
}

func GetGoal(c *gin.Context) {
	goalID := c.Param("id")

	var goal models.Goal
	results := initializers.DB.Where("id = ?", goalID).Preload("Milestones").First(&goal)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	goalResponse := helpers.CreateGoalResponse(goal)
	helpers.RespondWithSuccess(c, http.StatusOK, "Goal fetched successfully", "000", goalResponse)
}

func UpdateGoal(c *gin.Context) {
	goalID := c.Param("id")

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request", "001")
		return
	}

	if err := helpers.ValidateRequest(req, "Goal"); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "001")
		return
	}

	var goal models.Goal
	results := initializers.DB.Where("id = ?", goalID).First(&goal)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Record not found", "500")
		return
	}

	goal.Title = req["title"].(string)
	goal.Description = req["description"].(string)

	results = initializers.DB.Save(&goal)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	goalResponse := helpers.CreateGoalResponse(goal)
	helpers.RespondWithSuccess(c, http.StatusOK, "Goal updated successfully", "000", goalResponse)
}

func DeleteGoal(c *gin.Context) {
	goalID := c.Param("id")

	var goal models.Goal
	results := initializers.DB.Where("id = ?", goalID).Preload("Milestones").First(&goal)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	results = initializers.DB.Delete(&goal)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	helpers.RespondWithSuccess(c, http.StatusOK, "Goal deleted successfully", "000", nil)
}
