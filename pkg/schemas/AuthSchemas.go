package schemas

type RegistrationRequestSchema struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type LoginRequestSchema struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
