package utils

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type client struct {
	limiter *rate.Limiter
	lastSeen time.Time
}

var(
	Mu sync.Mutex
	Clients=make(map[string]*client)
)

func GetRateLimiter(ip string,limit rate.Limit , burst int)*rate.Limiter{
	Mu.Lock()
	defer Mu.Unlock()

	if c,exists:=Clients[ip];exists{
		c.lastSeen=time.Now()
		return c.limiter
	}

	limiter:=rate.NewLimiter(limit,burst)
	Clients[ip]=&client{
		limiter:limiter ,
		lastSeen: time.Now(),
	}

	return limiter
}

func AllowRequest(ip string, limit rate.Limit, burst int) (bool,time.Duration){
	limiter:=GetRateLimiter(ip,limit,burst)
	r:=limiter.Reserve()
	if !r.OK(){
		return false,0
	}

	delay:=r.Delay()
	if delay >0{
		r.Cancel()
		return false,delay
	}
	return true,0 
}