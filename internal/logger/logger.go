package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger(isDev bool){

	var err error
	if isDev{
		Log,err=zap.NewDevelopment()
		Log.Info("Logger Initialized (Development)")
	}else{
		Log,err=zap.NewProduction()
		Log.Info("Logger Initialized (Production)")
	}

	if err!=nil{
		panic(err)
	}
} 

func Sugar() *zap.SugaredLogger{
	return Log.Sugar()
}