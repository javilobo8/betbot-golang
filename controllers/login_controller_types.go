package controllers

type CreateUserRequestBody struct {
	Username string `json:"username" validate:"min=4,max=16,regexp=^[a-zA-Z0-9]*$"`
	Password string `json:"password" validate:"min=8,max=16,regexp=^[a-zA-Z0-9]*$"`
}
