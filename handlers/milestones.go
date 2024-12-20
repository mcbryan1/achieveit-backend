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
	results := initializers.DB.Where("goal_id = ?", uid).Preload("Comments").Find(&milestones)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	milestoneResponse := helpers.FetchMilestonesResponse(milestones)

	helpers.RespondWithSuccess(c, http.StatusOK, "Milestones fetched successfully", "000", milestoneResponse)
}

func DeleteMilestone(c *gin.Context) {
	// When deleting a milestone, all comments associated with the milestone should be deleted as well

	milestoneID := c.Param("id")

	if milestoneID == "" {
		helpers.RespondWithError(c, http.StatusBadRequest, "Milestone ID is required", "001")
		return
	}

	uid, err := uuid.Parse(milestoneID)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Invalid milestone ID", "500")
		return
	}

	var milestone models.Milestone
	results := initializers.DB.Where("id = ?", uid).Preload("Comments").First(&milestone)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	// Get the associated goal
	var goal models.Goal
	results = initializers.DB.Where("id = ?", milestone.GoalID).Preload("Milestones").First(&goal)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Error fetching goal", "500")
		return
	}

	// Delete the milestone
	results = initializers.DB.Delete(&milestone)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	// Recalculate the progress of the goal
	var completedMilestones int
	for _, milestone := range goal.Milestones {
		if milestone.ID != uid && milestone.Completed {
			completedMilestones++
		}
	}

	goal.Progress = (float64(completedMilestones) / float64(len(goal.Milestones)-1)) * 100

	results = initializers.DB.Save(&goal)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	helpers.RespondWithSuccess(c, http.StatusOK, "Milestone deleted successfully", "000", nil)
}

func UpdateMilestone(c *gin.Context) {
	mileStoneId := c.Param("id")

	if mileStoneId == "" {
		helpers.RespondWithError(c, http.StatusBadRequest, "Milestone ID is required", "001")
		return
	}

	uid, err := uuid.Parse(mileStoneId)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Invalid milestone ID", "500")
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request", "001")
		return
	}

	var milestone models.Milestone
	results := initializers.DB.Where("id = ?", uid).First(&milestone)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Record not found", "500")
		return
	}

	// Title can't be empty
	if title, ok := req["title"]; ok && title == "" {
		helpers.RespondWithError(c, http.StatusBadRequest, "Title cannot be empty", "001")
		return
	} else if ok {
		milestone.Title = title.(string)
	}

	// Logic for updating the completed status of the milestone
	if completed, ok := req["completed"]; ok {
		completedBool, ok := completed.(bool)
		if !ok {
			helpers.RespondWithError(c, http.StatusBadRequest, "Invalid value for completed", "001")
			return
		}
		milestone.Completed = completedBool

		var goal models.Goal
		results = initializers.DB.Where("id = ?", milestone.GoalID).Preload("Milestones").First(&goal)
		if results.Error != nil {
			helpers.RespondWithError(c, http.StatusInternalServerError, "Error fetching goal", "500")
			return
		}

		var completedMilestones int
		for _, milestone := range goal.Milestones {
			if milestone.Completed {
				completedMilestones++
			}
		}

		goal.Progress = (float64(completedMilestones) / float64(len(goal.Milestones))) * 100

		results = initializers.DB.Save(&goal)
		if results.Error != nil {
			helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
			return
		}
	}

	results = initializers.DB.Save(&milestone)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	milestoneResponse := helpers.CreateMilestoneResponse(milestone)
	helpers.RespondWithSuccess(c, http.StatusOK, "Milestone updated successfully", "000", milestoneResponse)
}
