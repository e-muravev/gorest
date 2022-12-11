package repositories

import (
	"app/dbstorage"
	"app/user/models"

	"gorm.io/gorm"
)

type UserCreateRepository struct {
	conn *gorm.DB
}

func NewUserCreateRepository() UserCreateRepository {
	return UserCreateRepository{conn: dbstorage.NewDB().Db}
}

func (u *UserCreateRepository) UserCreate(name string, email string, passowrd string) models.UserModel {

	user := models.UserModel{
		Name:     name,
		Email:    email,
		Password: passowrd,
		// Ctime:    time.Now(),
	}

	u.conn.Create(&user)
	return user
}
