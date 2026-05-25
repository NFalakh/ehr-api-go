package dto

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

// Response is the standard API response wrapper.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

// Meta contains pagination metadata.
type Meta struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// SuccessResponse sends a successful JSON response.
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// PaginatedResponse sends a paginated JSON response.
func PaginatedResponse(c *gin.Context, statusCode int, message string, data interface{}, total int64, page, perPage int) {
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta: &Meta{
			Page:       page,
			PerPage:    perPage,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}

// ErrorResponse sends an error JSON response.
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
	})
}

type PatientResponse struct {
	ID             uint      `json:"id"`
	IdentityNumber string    `json:"identity_number"`
	Name           string    `json:"name"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	BloodType      string    `json:"blood_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
