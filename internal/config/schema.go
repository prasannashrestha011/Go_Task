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
		Url string
	}
	Jwt struct {
		Secret        string
		AccessExpiry  time.Duration
		RefreshExpiry time.Duration 
	}
}