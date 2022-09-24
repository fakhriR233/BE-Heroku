package database

import (
	"dumbflix_be/models"
	"dumbflix_be/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
  err := mysql.DB.AutoMigrate(&models.User{},&models.Transaction{},&models.Category{},&models.Film{},&models.Episode{})

  if err != nil {
    fmt.Println(err)
    panic("Migration Failed")
  }

  fmt.Println("Migration Success")
}