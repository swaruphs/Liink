package models

type User struct {
	Id       int    `json:"id",schema:"-"`
	Name     string `json:"username",schema:"username"`
	Password string `json:"-",schema:"password"` // not returning password
}
