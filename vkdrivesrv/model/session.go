package model

import "time"

type Session struct {
	UserId    int
	Token     string
	UserAgent string
	Date      time.Time `sql:"DEFAULT:now()"`
}



