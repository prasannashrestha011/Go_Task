package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger(isDev bool){

	var err error
	if isDev{
		Log,err=zap.NewDevelopment(zap.AddStacktrace(zap.DPanicLevel))
		Log.Info("Logger Initialized (Development)")
	}else{
		Log,err=zap.NewProduction(zap.AddStacktrace(zap.DPanicLevel))
		Log.Info("Logger Initialized (Production)")
	}

	if err!=nil{
		panic(err)
	}
} 

func Sugar() *zap.SugaredLogger{
	return Log.Sugar()
}