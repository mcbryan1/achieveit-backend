package helpers

import (
	"github.com/mcbryan1/achieveit-backend/initializers"
	"github.com/mcbryan1/achieveit-backend/models"
)

func CreateUserResponse(user models.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"created_at": user.CreatedAt,
	}
}

func CreateGoalResponse(goal models.Goal) map[string]interface{} {
	return map[string]interface{}{
		"id":         goal.ID,
		"title":      goal.Title,
		"decription": goal.Description,
		"progress":   goal.Progress,
		"created_at": goal.CreatedAt,
		"milestones": FetchMilestonesResponse(goal.Milestones),
	}
}

func FetchGoalsResponse(goals []models.Goal) []map[string]interface{} {
	var responseArray []map[string]interface{}
	for _, goal := range goals {
		// Recalculate progress
		var completedMilestones int
		for _, milestone := range goal.Milestones {
			if milestone.Completed {
				completedMilestones++
			}
		}
		goal.Progress = (float64(completedMilestones) / float64(len(goal.Milestones))) * 100

		// Save the updated progress
		initializers.DB.Save(&goal)

		response := map[string]interface{}{
			"id":          goal.ID,
			"title":       goal.Title,
			"description": goal.Description,
			"progress":    goal.Progress,
			"created_at":  goal.CreatedAt,
			"milestones":  FetchMilestonesResponse(goal.Milestones),
		}
		responseArray = append(responseArray, response)
	}
	return responseArray
}

func CreateMilestoneResponse(milestone models.Milestone) map[string]interface{} {
	return map[string]interface{}{
		"id":        milestone.ID,
		"title":     milestone.Title,
		"completed": milestone.Completed,
	}
}

func FetchMilestonesResponse(milestones []models.Milestone) []map[string]interface{} {
	var responseArray []map[string]interface{}
	for _, milestone := range milestones {
		response := map[string]interface{}{
			"id":        milestone.ID,
			"title":     milestone.Title,
			"completed": milestone.Completed,
			"comments":  FetchCommentsResponse(milestone.Comments),
		}
		responseArray = append(responseArray, response)
	}
	return responseArray
}

func CreateCommentResponse(comment models.Comment) map[string]interface{} {
	return map[string]interface{}{
		"id":         comment.ID,
		"content":    comment.Content,
		"created_at": comment.CreatedAt,
	}
}

func FetchCommentsResponse(comments []models.Comment) []map[string]interface{} {
	var responseArray []map[string]interface{}
	for _, comment := range comments {
		response := map[string]interface{}{
			"id":         comment.ID,
			"content":    comment.Content,
			"created_at": comment.CreatedAt,
		}
		responseArray = append(responseArray, response)
	}
	return responseArray
}
