package config

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func getGormPostgresUrl(dbConfig DBConfiguration) string {
	return fmt.Sprintf(
		"host=%s user=%s port=%d dbname=%s sslmode=disable password=%s",
		dbConfig.DbHost,
		dbConfig.DbUser,
		dbConfig.DbPort,
		dbConfig.DbName,
		dbConfig.DbPass,
	)
}

func CreateDB(conf DBConfiguration) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(getGormPostgresUrl(conf)), &gorm.Config{})
}