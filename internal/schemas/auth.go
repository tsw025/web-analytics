package schemas

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,lowercase,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32,alphanum,lowercase"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}
