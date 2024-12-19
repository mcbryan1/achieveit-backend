package helpers

import "github.com/mcbryan1/achieveit-backend/models"

func CreateUserResponse(user models.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"created_at": user.CreatedAt,
	}
}
