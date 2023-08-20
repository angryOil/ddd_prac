package model

type LoginModel struct {
	Email string `json:"email"`
	Pw    string `json:"pw"`
}

type CreateUserModel struct {
	Email string `json:"email"`
	Pw    string `json:"pw"`
}
