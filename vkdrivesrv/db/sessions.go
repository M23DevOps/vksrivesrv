package db

import (
	_ "log"
	_ "fmt"
	_ "errors"

	_ "gopkg.in/gormigrate.v1"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"../model"
)

//SelectSession select user's session
func SelectSession(strToken, strUserAgent string) (model.Session, error) {
	db := connect()

	var session model.Session
	res := db.Where("token = ? and user_agent = ? and date > now() - interval '24 hours'", strToken, strUserAgent).First(&session)

	return session, res.Error
}

//InsertSession insert user's session
func InsertSession(intUserId int, strToken, strUserAgent string) (error) {
	db := connect()

	session := &model.Session{
		UserId:    intUserId,
		Token:     strToken,
		UserAgent: strUserAgent,
	}

	err := db.Create(&session)

	return err.Error
}

//DeleteSession delete user's session
func DeleteSession(strToken, strUserAgent string) (error) {
	db := connect()

	err := db.Where("token = ? and user_agent = ?", strToken, strUserAgent).Delete(model.Session{})
	return err.Error
}