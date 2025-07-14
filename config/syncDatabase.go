package config

import (
	"fmt"

	model "main.go/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&model.UserModel{})
	if err != nil {
		fmt.Println("Failed to migrate model")
		return
	}

}
