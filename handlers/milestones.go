package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mcbryan1/achieveit-backend/helpers"
	"github.com/mcbryan1/achieveit-backend/initializers"
	"github.com/mcbryan1/achieveit-backend/models"
)

func CreateMilestone(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request", "001")
		return
	}

	if err := helpers.ValidateRequest(req, "Milestone"); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "001")
		return
	}

	// Create the milestone
	newMilestone := models.Milestone{
		Title:     req["title"].(string),
		Completed: false,
		GoalID:    uuid.MustParse(req["goal_id"].(string)),
	}

	results := initializers.DB.Create(&newMilestone)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	milestoneResponse := helpers.CreateMilestoneResponse(newMilestone)

	helpers.RespondWithSuccess(c, http.StatusOK, "Milestone created successfully", "000", milestoneResponse)
}

func GetMilestones(c *gin.Context) {
	// Should be filtered by goal_id
	// Get the goal_id from query params(required)

	goalID := c.Query("goal_id")

	if goalID == "" {
		helpers.RespondWithError(c, http.StatusBadRequest, "goal_id is required", "001")
		return
	}

	uid, err := uuid.Parse(goalID)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Invalid goal ID", "500")
		return
	}

	var milestones []models.Milestone
	results := initializers.DB.Where("goal_id = ?", uid).Find(&milestones)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	milestoneResponse := helpers.FetchMilestonesResponse(milestones)

	helpers.RespondWithSuccess(c, http.StatusOK, "Milestones fetched successfully", "000", milestoneResponse)
}
