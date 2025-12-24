package services

import (
	"context"
	"main/internal/repository"
	"main/internal/schema"
	"main/internal/utils"
	"net/http"
)

type AuthService interface {
	Login(ctx context.Context, creds *schema.UserLoginDTO) (*schema.LoginMetaDataDTO, error)
	VerifyEmail(ctx context.Context, creds *schema.VerifyEmailRequest) (bool,string, error)
}

type authService struct {
	repo repository.UserRepository
}


func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (a *authService) Login(ctx context.Context, creds *schema.UserLoginDTO) (*schema.LoginMetaDataDTO, error) {
	email := creds.Email

	userDetails, err := a.repo.GetByEmail(ctx, email)

	if err != nil {
		err := utils.NewAppError(404, "INVALID_CREDENTIALS", "Email not found in the database", nil)
		return nil, err
	}
	isMatches := utils.ComparePassword(userDetails.Password, creds.Password)

	if !isMatches {
		err := utils.NewAppError(401, "INVALID_CREDENTIALS", "Invalid username or password", nil)
		return nil, err
	}

	userMetaData := &schema.LoginMetaDataDTO{
		ID:    userDetails.ID,
		Name:  userDetails.Name,
		Email: userDetails.Email,
	}

	return userMetaData, nil
}


func (a *authService) VerifyEmail(ctx context.Context, creds *schema.VerifyEmailRequest) (bool,string, error) {

	isVerified:=utils.VerifyCode(ctx,creds.Email,creds.Code)
	if !isVerified{
		err:=utils.NewAppError(http.StatusUnauthorized,"REQ_UNAUTHORIZED","Verification code didn't matched, please try again",nil)
		return false,"Verification code didnt matched, please try again",err
	}

	user, _ := a.repo.GetByEmail(ctx,creds.Email)
	
	
	utils.DeleteVerificationCode(ctx,creds.Email)
	user.IsVerified=true
	_, err := a.repo.Update(ctx,user)
	if err != nil {
		return false,"Failed to update user's verified status", err
	}

	return true,"Email address has been verified, you can now login",nil
}