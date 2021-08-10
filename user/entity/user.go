package entity

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
}
