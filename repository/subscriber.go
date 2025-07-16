package repository

import (
	"subscribers/model"

	"gorm.io/gorm"
)

func CreateNewUser(db *gorm.DB, user model.User) error {
	return db.Create(&user).Error
}

func UpdateUser(db *gorm.DB, user model.User) error {
	return db.Save(user).Error
}

func GetUserById(db *gorm.DB, id string) (model.User, error) {
	var user model.User
	result := db.First(&user, "id = ?", id)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, result.Error
}

func CalculateSubsCost(db *gorm.DB, userId string, serviceName string) (int, error) {
	var result int

	q := db.Model(&model.User{}).Select("SUM(price) as total").Where("user_id	 = ? AND service_name = ?", userId, serviceName).Scan(&result)

	return result, q.Error
}

func GetUsers(db *gorm.DB, users *[]model.User) error {
	return db.Find(&users).Error
}

func DeleteUser(db *gorm.DB, id string) error {
	return db.Delete(&model.User{}, "id = ?", id).Error
}
