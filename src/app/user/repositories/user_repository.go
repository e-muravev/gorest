package repositories

import (
	"app/dbstorage"
	"app/user/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository() UserCreateRepository {
	return UserCreateRepository{conn: dbstorage.NewDB().Db}
}

func (u *UserCreateRepository) UserExistsByName(name string) (models.UserModel, error) {
	user := models.UserModel{}

	result := u.conn.Where("name = ?", name).Take(&user)
	fmt.Println(result.Error)

	return user, result.Error
}
