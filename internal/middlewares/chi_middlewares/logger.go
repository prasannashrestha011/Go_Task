package chimiddlewares

import (
	"main/internal/logger"
	"net/http"
	"time"

	"go.uber.org/zap"
)

/*
Unlike gin, chi doesnot have builtin logger middleware,
So I have implemented a simple logger system using zap
*/

func LoggerMiddleware(next http.Handler)http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start:=time.Now()
		logger.Log.Info("Incoming request: ",zap.String("method",r.Method),
											zap.String("path",r.URL.Path),
											zap.String("remote address",r.RemoteAddr))

		next.ServeHTTP(w,r)
		duration:=time.Since(start)

		logger.Log.Info("Request completed",zap.String("method",r.Method),
											zap.String("path",r.URL.Path),
											zap.String("duration",duration.String()))														
	})
}