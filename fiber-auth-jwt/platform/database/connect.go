// ./connect.go
package database

import (
	"fmt"
	"strconv"

	"github.com/afrianska/go-try/fiber-auth-jwt/app/models"
	"github.com/afrianska/go-try/fiber-auth-jwt/pkg/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB()  {
	var err error

	p := configs.EnvConf("DB_PORT")
	port , err := strconv.ParseUint(p, 10, 30)
	if err != nil {
		panic("Failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", configs.EnvConf("DB_HOST"), configs.EnvConf("DB_USER"), configs.EnvConf("DB_PASS"), configs.EnvConf("DB_NAME"), port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Connection opened to database")

	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)
	fmt.Println("Database migrated")
}