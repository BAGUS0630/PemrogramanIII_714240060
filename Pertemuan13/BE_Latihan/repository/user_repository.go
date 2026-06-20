package repository

import (
	"BE_LATIHAN/config"
	"BE_LATIHAN/model"
)

func FindUserByUsername(username string) (model.User, error) {
	var user model.User
	result := config.GetDB().First(&user, "username = ?", username)
	return user, result.Error
}

func InsertUser(user *model.User) (*model.User, error) {
	result := config.GetDB().Create(user)
	return user, result.Error
}

func UpdateUserPassword(username string, newHashedPassword string) error {
	return config.GetDB().Model(&model.User{}).Where("username = ?", username).Update("password", newHashedPassword).Error
}