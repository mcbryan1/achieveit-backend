package helpers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mcbryan1/achieveit-backend/initializers"
	"github.com/mcbryan1/achieveit-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(username string) (models.User, error) {
	var user models.User
	err := initializers.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func UserExists(username string) bool {
	var user models.User
	result := initializers.DB.Where("username = ?", username).First(&user)
	return result.Error == nil
}

func GenerateJWTToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func RespondWithSuccess(c *gin.Context, code int, message interface{}, respCode string, data ...interface{}) {
	response := struct {
		RespCode string      `json:"resp_code"`
		RespDesc interface{} `json:"resp_desc"`
		Data     interface{} `json:"data"`
	}{
		RespCode: respCode,
		RespDesc: message,
		Data:     nil,
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	c.JSON(code, response)
}

func RespondWithError(c *gin.Context, code int, message interface{}, resCode string) {
	c.AbortWithStatusJSON(code, gin.H{"resp_desc": message, "resp_code": resCode})
}

func ParseRequest(c *gin.Context) (map[string]interface{}, error) {
	var req map[string]interface{}
	err := c.ShouldBindJSON(&req)
	return req, err
}

func CheckPassword(user models.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func ProcessLogin(c *gin.Context) (req map[string]interface{}, user models.User, tokenString string, err error) {
	req, err = ParseRequest(c)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "Invalid request data", "001")
	}

	user, err = GetUser(req["username"].(string))
	if err != nil {
		RespondWithError(c, http.StatusUnauthorized, "Invalid username or password", "001")
		return
	}

	if err = CheckPassword(user, req["password"].(string)); err != nil {
		RespondWithError(c, http.StatusUnauthorized, "Invalid username or password", "001")
		return
	}

	tokenString, err = GenerateJWTToken(user)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, "Could not generate token", "500")
		return
	}

	return
}

func ValidateRequest(req map[string]interface{}, req_type string) error {
	var requiredFields []string

	switch req_type {
	case "User":
		requiredFields = []string{"username", "password"}
	// case "Mock":
	// 	requiredFields = []string{"enrollment", "gh_language", "french", "school_name", "contact", "location"}
	// case "Terminal":
	// 	requiredFields = []string{"nursery_enrollment", "kg_enrollment", "basic_1_enrollment", "basic_2_enrollment", "basic_3_enrollment", "basic_4_enrollment", "basic_5_enrollment", "basic_6_enrollment", "basic_7_enrollment", "basic_8_enrollment", "school_name", "contact", "location"}
	default:
		return fmt.Errorf("invalid request type")
	}

	for _, field := range requiredFields {
		if _, ok := req[field]; !ok {
			return fmt.Errorf("%s is required", field)
		}
		// Trim whitespace from the field value if it's a string
		if strVal, ok := req[field].(string); ok {
			strVal = strings.TrimSpace(strVal)
			if strVal == "" {
				return fmt.Errorf("%s cannot be empty", field)
			}
			req[field] = strVal
		}
	}

	// Additional validation for User type
	// if req_type == "User" {

	// 	phoneNumber, ok := req["phone_number"].(string)
	// 	if !ok || !IsValidPhoneNumber(phoneNumber) {
	// 		return fmt.Errorf("invalid phone number")
	// 	}
	// }

	return nil
}

// func IsValidPhoneNumber(phoneNumber string) bool {
// 	if len(phoneNumber) != 10 {
// 		return false
// 	}

// 	if !regexp.MustCompile(`^\d+$`).MatchString(phoneNumber) {
// 		return false
// 	}

// 	return true
// }
