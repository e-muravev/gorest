package actions

import (
	"app/user"
)

type UserCreateAction struct{}

func (a UserCreateAction) Run(serializerData user.transformers.UserCreateTransformer) (models.UserModel, error) {
	var user user.models.UserModel

	if !handlers.PasswordMinLengthValidator(&serializerData.Password) {
		return user, user.exceptions.MinLengthInvalidPasswordError
		// return user, &exceptions.MinLengthInvalidPasswordError{
		// 	Title:  "password",
		// 	Detail: "enter a valid password",
		// 	// Err:    errors.New("enter_a_valid_password"),
		// }
	}

	if !handlers.PasswordComparingValidator(&serializerData.Password, &serializerData.PasswordRepeated) {
		return user, user.exceptions.PasswordComparingError
	}

	// Check if user exists
	if handlers.UserExistsByEmailHandler(serializerData.Email) {
		return user, user.exceptions.UserExistsByEmailError
	}

	_, err := handlers.UserExistsByNameHandler(serializerData.Name)
	// fmt.Println("\n", user_exists, err, "\n")
	if err == nil {
		return user, user.exceptions.UserExistsByNameError
	}


	serializerData.Password, _ = user.handlers.PasswordHashingHandler(serializerData.Password)

	user = user.handlers.CreateNewUserHandler{Repository: user.repositories.NewUserCreateRepository()}.Run(
		serializerData.Name,
		serializerData.Email,
		serializerData.Password,
	)

	return user, nil
}
