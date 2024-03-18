package db

import (
	"log"

	"github.com/sgokul961/echo-hub-notification-svc/pkg/config"
	"github.com/sgokul961/echo-hub-notification-svc/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(c config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.DBUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)

	}
	db.AutoMigrate(&domain.Notification{})

	return db, err
}
