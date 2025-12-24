package config

import "time"

type AppConfigs struct {
	Server struct {
		Port struct{
			Auth string
			User string
			Order string
		}
		Env string
	}
	Database struct {
		Postgres string
		Redis string
	}
	Jwt struct {
		Secret        string
		AccessExpiry  time.Duration
		RefreshExpiry time.Duration 
	}
	Resend struct{
		ApiKey string 
		AppDomain string
	}
}