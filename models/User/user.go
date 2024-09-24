package models

import (
	"ShushiNaritai/utils"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// 创建用户
func CreateUser(user *User) error {
	result := utils.DB.Create(user)
	return result.Error
}

// 查找用户
func GetUserByID(id uint) (*User, error) {
	var user User
	result := utils.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// 更新用户
func UpdateUser(user *User) error {
	result := utils.DB.Save(user)
	return result.Error
}

// 删除用户
func DeleteUser(id uint) error {
	result := utils.DB.Delete(&User{}, id)
	return result.Error
}
