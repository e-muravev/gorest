package transformers

type UserCreateTransformer struct {
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email" binding:"required"`
	Password         string `json:"password" binding:"required"`
	PasswordRepeated string `json:"password_repeated" binding:"required"`
}
