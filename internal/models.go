package internal

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

type ResetPassForm struct {
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

type Car struct {
	ID    int
	Name  string
	Make  string
	Model string
	Year  int
}

type CarForm struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Make  string `json:"make" xml:"make" form:"make" query:"make"`
	Model string `json:"model" xml:"model" form:"model" query:"model"`
	Year  int    `json:"year" xml:"year" form:"year" query:"year"`
}
