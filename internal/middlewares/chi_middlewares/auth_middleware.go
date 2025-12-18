package chimiddlewares

import (
	"context"
	"main/internal/logger"
	"main/internal/utils"
	"net/http"
	"strings"

	"go.uber.org/zap"
)
type contextKey string
const userKEY=contextKey("userID")

func JWTAuthMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader:=r.Header.Get("Authorization")
		parts:=strings.Split(authHeader, "Bearer ")
		if len(parts)!=2 || strings.TrimSpace(parts[1])==""{
			http.Error(w,"Request Unauthorized, please login to your session",http.StatusUnauthorized)
			return
		} 
		tokenStr:=parts[1]
		token,err:=utils.ValidateJWT(tokenStr)
		if err!=nil{
			logger.Log.Info("JWT validation error",zap.Error(err))
			http.Error(w,"Session expired, refresh your access token",http.StatusUnauthorized)
			return
		}
		userID,err:=utils.GenerateUserIDFromToken(token)
		if err!=nil{
			logger.Log.Info("JWT ID extraction error",zap.Error(err))
			http.Error(w,"Invalid authentication token, please login again!!",http.StatusUnauthorized)
			return
		}
		ctx:=context.WithValue(r.Context(),userKEY,userID)
		r=r.WithContext(ctx)
		next.ServeHTTP(w,r)
	})
}