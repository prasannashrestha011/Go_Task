package services

import (
	"log"
	"main/internal/models"
	"main/internal/repository"
	"main/internal/utils"

	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(user *models.UserCreateDTO) (*models.UserResponseDTO, error)
	GetUserByID(id uuid.UUID) (*models.UserResponseDTO, error)
	GetUsers() ([]*models.UserResponseDTO, error)
	GetUserByEmail(email string) (*models.UserResponseDTO, error)
	UpdateUser(id uuid.UUID, user *models.UserUpdateDTO) (*models.UserResponseDTO, error)
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
}

func (u *userService) GetUsers() ([]*models.UserResponseDTO, error) {
	users,err:=u.repo.GetAll()
	if err!=nil{
		return nil,err
	}

	var usersDTO []*models.UserResponseDTO
	for _,user:=range users{
		dto:=&models.UserResponseDTO{
			ID:user.ID ,
			Name: user.Name,
			Email: user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		usersDTO = append(usersDTO,dto )
	}
	return usersDTO,nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) RegisterUser(dto *models.UserCreateDTO) (*models.UserResponseDTO, error) {
	user := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
	err := u.repo.Create(user)
	if err != nil {
		return nil, err
	}
	userDTO := utils.ToUserResponseDTO(user)
	return userDTO, nil
}

func (u *userService) DeleteUser(id uuid.UUID) error {
	return u.repo.Delete(id)
}

func (u *userService) UpdateUser(id uuid.UUID, dto *models.UserUpdateDTO) (*models.UserResponseDTO, error) {

	user, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	log.Println(dto)
	if dto.Name != "" {
		user.Name = dto.Name
	}

	if dto.Email != "" {
		user.Email = dto.Email
	}
	if dto.Password != "" {
		user.Password = dto.Password
	}

	updated_user, err := u.repo.Update(user)
	if err != nil {
		return nil, err
	}
	userDTO := utils.ToUserResponseDTO(updated_user)
	return userDTO, nil
}

func (u *userService) GetUserByEmail(email string) (*models.UserResponseDTO, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	userDTO := utils.ToUserResponseDTO(user)
	return userDTO, nil
}

func (u *userService) GetUserByID(id uuid.UUID) (*models.UserResponseDTO, error) {
	user, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	userDTO := utils.ToUserResponseDTO(user)
	return userDTO, nil
}
