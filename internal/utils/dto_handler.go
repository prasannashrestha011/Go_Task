// Methods to convert database model into DTO objects for response schema
package utils

import "main/internal/models"

func ToUserResponseDTO(user *models.User)(*models.UserResponseDTO){
	return &models.UserResponseDTO{
				ID: user.ID,
				Name: user.Name,
				Email: user.Email,
				CreatedAt: user.CreatedAt,
			}
}