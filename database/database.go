package database

import (
	"log"

	"github.com/playsoil/todo-go/models"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB *gorm.DB
}

var DB DBInstance

func ConnectDB(db *gorm.DB) {

	log.Println("connected to db. running migrations...")

	db.AutoMigrate(&models.Task{})

	DB = DBInstance{DB: db}
}
