package app

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest_api/utils"
)

func GetDBEngine() (*gorm.DB, error) {
	conf := utils.New()
	dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		conf.PostgreSQL.PostgreHost,
		conf.PostgreSQL.PostgrePort,
		conf.PostgreSQL.PostgreUser,
		conf.PostgreSQL.PostgreDbName,
		conf.PostgreSQL.PostgrePassword,
	)
	engine, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return engine, err
	}
	return engine, nil
}
