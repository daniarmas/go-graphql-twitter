package gographqltwitter

import "time"

type UserRepo interface {
	
}

type User struct {
	ID         string
	Username   string
	Email      string
	Password   string
	CreateTime time.Time
	UpdateTime time.Time
}
