package requests

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterBody struct {
	Name            string `json:"name" binding:"required,min=4,max=10"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6,max=16,eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type ForgotPasswordBody struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordBody struct {
	Password string `json:"password" binding:"required"`
}
