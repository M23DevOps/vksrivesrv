package db

import (
	_ "log"
	_ "time"
	_ "fmt"
	_ "errors"

	_ "gopkg.in/gormigrate.v1"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"../model"
)

//SelectUser select user by id
func SelectUser(intId int) (model.User, error) {
	db := connect()

	var user model.User
	res := db.Where("id = ?", intId).First(&user)

	return user, res.Error
}

//SelectUserBySId select user by social id
func SelectUserBySNId(intSNId int, strSNType string) (model.User, error) {
	db := connect()

	var user model.User
	res := db.Where("sn_id = ? and sn_type = ?", intSNId, strSNType).First(&user)

	return user, res.Error
}

//UpdateUser update user
func UpdateUser(intId int, strFirstName, strLastName, strPhotoURL string, bIsActive bool) (error) {
	db := connect()

	err := db.Table("users").Where("id = ?", intId).Updates(model.User{FirstName: strFirstName, LastName: strLastName, PhotoURL: strPhotoURL})
	err = db.Table("users").Where("id = ?", intId).Update("is_active", bIsActive)
	return err.Error
}

//UpdateUserStatus update user status: available or not
func UpdateUserStatus(intId int, bIsActive bool) (error) {
	db := connect()

	err := db.Table("users").Where("id = ?", intId).Update("is_active", bIsActive)
	return err.Error
}

//InsertUser insert user
func InsertUser(intSNId int, strSNType, strFirstName, strLastName, strPhotoURL string) (error) {
	db := connect()

	user := &model.User{
		//Id:        nil,
		SNType:    strSNType,
		SNId:      intSNId,
		FirstName: strFirstName,
		LastName:  strLastName,
		PhotoURL:  strPhotoURL,
		IsActive:  true,
	}

	err := db.Create(&user)

	return err.Error
}