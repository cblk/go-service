package models

type User struct {
	Username string `form:"username" json:"username" description:"The user's username"`
	Password string `form:"password" json:"password" description:"The user's password"`
	Name     string `form:"name" json:"name" description:"The user's real name"`
}
