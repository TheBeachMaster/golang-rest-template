package models

type UserInfo struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
	UserID       string `json:"id"`
}

type (
	CreateUser struct {
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
		Username     string `json:"userName"`
		EmailAddress string `json:"emailAddress"`
		Password     string `json:"password"`
	}
)
