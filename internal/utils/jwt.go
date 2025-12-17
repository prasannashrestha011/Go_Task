package utils

import (
	"main/internal/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var jwtSecret = []byte("Secret_Key")

func GenerateTokens(userID uuid.UUID)(accessToken string,refreshToken string,err error){
	atClaims:=jwt.MapClaims{
		"sub":userID.String(),
		"exp":time.Now().Add(15 * time.Minute).Unix(),
		"iat":time.Now().Unix(),
	}

	at:=jwt.NewWithClaims(jwt.SigningMethodHS256,atClaims)
	accessToken,err=at.SignedString(jwtSecret)
	if err!=nil{
		logger.Log.Error("JWT error:",zap.Error(err))
		return "","",err
	}

	rtClaims:=jwt.MapClaims{
		"sub":userID.String(),
		"exp":time.Now().Add(7*24*time.Hour).Unix(),
		"iat":time.Now().Unix(),
	}
	rt:=jwt.NewWithClaims(jwt.SigningMethodHS256,rtClaims)
	refreshToken,err=rt.SignedString(jwtSecret)
	if err!=nil{
		logger.Log.Error("JWT error:",zap.Error(err))
		return "","",err
	}

	return accessToken,refreshToken,nil
}

func ValidateJWT(tokenString string)(*jwt.Token,error){
	token,err:=jwt.Parse(tokenString,func(t *jwt.Token) (any, error) {
		if _,ok:=t.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,jwt.ErrInvalidType
		}
		return jwtSecret,nil
	})
	if err!=nil{
		return nil,err
	}
	if !token.Valid{
		return nil,jwt.ErrTokenInvalidClaims
	}
	return token,nil
}

func GenerateUserIDFromToken(token *jwt.Token)(string,error){
	claims,ok:=token.Claims.(jwt.MapClaims)
	if !ok{
		return "",jwt.ErrTokenInvalidClaims
	}
	sub,ok:=claims["sub"].(string)
	if !ok{
		return "",jwt.ErrTokenInvalidClaims
	}
	return sub,nil
}