package model

type BaseUser struct {
	UserId     uint
	OpenId     string
	UserName   string
	Tel        string
	Email      string
	Permission int
	LastTime   string
}
