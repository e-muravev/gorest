package transformers

import "time"

type UserCreateRespTransformer struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdat"`
}
