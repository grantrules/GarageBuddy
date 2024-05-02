package main

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type LoginForm struct {
	Email    string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}

type RegisterForm struct {
	Name            string `json:"name" xml:"name" form:"name" query:"name"`
	Email           string `json:"email" xml:"email" form:"email" query:"email"`
	Password        string `json:"password" xml:"password" form:"password" query:"password"`
	PasswordConfirm string `json:"password-confirm" xml:"password-confirm" form:"password-confirm" query:"password-confirm"`
}
