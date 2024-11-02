package configuration

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = InitializeDB()

	if err != nil {
		return fmt.Errorf("erro initializing sqlite: %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initializer logger
	logger := NewLogger(p)
	return logger
}
