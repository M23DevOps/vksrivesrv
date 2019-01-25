package db

import (
	_ "log"
	_ "time"
	_ "fmt"
	_ "errors"

	"../model"
)

//SelectImagesByUser select all users's images by user id
func SelectImagesByUser(intId int) ([]model.Image, error) {
	db := connect()

	images := []model.Image{}
	res := db.Where("user_id = ? and is_delete = 0", intId).Find(&images)

	return images, res.Error
}

//UpdateImage by id and user id
func UpdateImage(intId, intUserId, intFolderId, intIsDelete int, strLabel  string) (error) {
	db := connect()

	err := db.Table("images").Where("id = ? and user_id = ?", intId, intUserId).Updates(model.Image{FolderId: intFolderId, Label: strLabel, IsDelete: intIsDelete})
	return err.Error
}

//InsertImage insert image
func InsertImage(intUserId, intFolderId int, strURL, strLabel string) (error) {
	db := connect()

	image := &model.Image{
		//Id:        nil,
		UserId: 	intUserId,
		URL:      	strURL,
		FolderId:	intFolderId,
		Label:  	strLabel,
		IsDelete:  	0,
	}

	err := db.Create(&image)

	return err.Error
}
