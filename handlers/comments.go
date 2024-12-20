package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mcbryan1/achieveit-backend/helpers"
	"github.com/mcbryan1/achieveit-backend/initializers"
	"github.com/mcbryan1/achieveit-backend/models"
)

func CreateComment(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request", "001")
		return
	}

	if err := helpers.ValidateRequest(req, "Comment"); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "001")
		return
	}

	// Create the comment
	newComment := models.Comment{
		Content:     req["content"].(string),
		MilestoneID: uuid.MustParse(req["milestone_id"].(string)),
	}

	results := initializers.DB.Create(&newComment)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	commentResponse := helpers.CreateCommentResponse(newComment)

	helpers.RespondWithSuccess(c, http.StatusOK, "Comment created successfully", "000", commentResponse)
}

func DeleteComment(c *gin.Context) {
	commentID := c.Param("id")
	comment := models.Comment{}

	results := initializers.DB.First(&comment, commentID)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusNotFound, "Comment not found", "404")
		return
	}

	results = initializers.DB.Delete(&comment)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	helpers.RespondWithSuccess(c, http.StatusOK, "Comment deleted successfully", "000", nil)
}

func UpdateComment(c *gin.Context) {
	commentId := c.Param("id")

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request", "001")
		return
	}

	var comment models.Comment
	results := initializers.DB.Where("id = ?", commentId).First(&comment)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Record not found", "500")
		return
	}

	content, ok := req["content"].(string)
	if !ok || content == "" {
		helpers.RespondWithError(c, http.StatusBadRequest, "Content cannot be empty", "001")
		return
	}

	comment.Content = content

	results = initializers.DB.Save(&comment)
	if results.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, results.Error, "500")
		return
	}

	commentResponse := helpers.CreateCommentResponse(comment)
	helpers.RespondWithSuccess(c, http.StatusOK, "Comment updated successfully", "000", commentResponse)
}
