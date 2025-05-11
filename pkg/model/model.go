package model

import "time"

type LogInfo struct {
	IpAddress    string
	User         string
	TimeStamp    time.Time
	RequestType  string
	RequestUrl   string
	ResponseCode int
	ResponseSize int64
	Platform     string
}
