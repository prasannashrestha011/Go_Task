package services

import (
	"context"
	"main/internal/repository"
	"main/internal/schema"
	"main/internal/utils"
)

type AuthService interface {
	Login(ctx context.Context, creds *schema.UserLoginDTO) (*schema.LoginMetaDataDTO, error)
}

type authService struct {
	repo repository.UserRepository
}


func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}


func (a *authService) Login(ctx context.Context, creds *schema.UserLoginDTO) (*schema.LoginMetaDataDTO, error) {
	email:=creds.Email

	userDetails,err:=a.repo.GetByEmail(ctx,email)

	if err!=nil{
		err:=utils.NewAppError(404,"INVALID_CREDENTIALS","Email not found in the database",nil)
		return nil,err
	}
	isMatches:=utils.ComparePassword(userDetails.Password,creds.Password)

	if !isMatches{
		err:=utils.NewAppError(401,"INVALID_CREDENTIALS","Invalid username or password",nil)
		return nil,err
	}

	userMetaData:=&schema.LoginMetaDataDTO{
		ID: userDetails.ID,
		Name: userDetails.Name,
		Email: userDetails.Email,
	}

	return userMetaData,nil
}