package db

import (
	"log"
	"time"

	gormigrate "gopkg.in/gormigrate.v1"

	_ "../config"
	// "../model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Migrate func for migrate
func Migrate() error {
	db := connect()
	defer db.Close()

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "201708161230",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					Id        int    `sql:"AUTO_INCREMENT" gorm:"primary_key"`
					SNType    string `sql:"not null"`
					SNId      int    `sql:"not null"`
					FirstName string `sql:"not null"`
					LastName  string `sql:"not null"`
					PhotoURL  string `sql:"not null"`
					IsActive  bool   `sql:"not null"`
				}
				return tx.AutoMigrate(&User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
		{
			ID: "201708161235",
			Migrate: func(tx *gorm.DB) error {
				type Session struct {
					UserId    int       `sql:"type:int REFERENCES users(id)"`
					Token     string    `sql:"not null"`
					UserAgent string    `sql:"not null"`
					Date      time.Time `sql:"DEFAULT:now()"`
				}
				return tx.AutoMigrate(&Session{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("sessions").Error
			},
		},
		{
			ID: "201709181245",
			Migrate: func(tx *gorm.DB) error {
				type UserConfig struct {
					UserId    int    `sql:"type:int REFERENCES users(id)"`
					UserAgent string `sql:"not null"`
					IsLogin   int    `sql:"not null"`
				}
				return tx.AutoMigrate(&UserConfig{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("user_configs").Error
			},
		},
		{
			ID: "201709181315",
			Migrate: func(tx *gorm.DB) error {
				type Image struct {
					Id			int		  	`sql:"AUTO_INCREMENT" gorm:"primary_key"`		
					UserId		int			`sql:"type:int REFERENCES users(id)"`
					URL			string      `sql:"not null"`
					Date		time.Time	`sql:"DEFAULT:now()"`
					FolderId	int
					Label		string
					IsDelete	int			`sql:"DEFAULT:0; not null"`
				}
				return tx.AutoMigrate(&Image{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("images").Error
			},
		},
	})
	err := m.Migrate()
	if err == nil {
		log.Printf("Migration did run successfully")
	} else {
		log.Printf("Could not migrate: %v", err)
	}

	return nil
}
