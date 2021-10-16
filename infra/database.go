package infra

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
)

func GetDbConn(DbHost string, DbPort string, DbName string, DbUser string, DbPassword string) (*gorm.DB, error) {
	connectionString := GetConnString("mysql", DbHost, DbPort, DbName, DbUser, DbPassword)
	dialector := GetDialector("mysql", connectionString)
	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Database connection failed %s", err)
	}

	return conn, nil
}

func GetDialector(db string, connString string) gorm.Dialector {
	if db == "postgresql" {
		return postgres.Open(connString)
	}

	return mysql.Open(connString)
}

func GetConnString(db string, DbHost string, DbPort string, DbName string, DbUser string, DbPassword string) string {
	if db == "postgresql" {
		return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
			DbHost, DbPort, DbName, DbUser, DbPassword)
	}
	
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	DbUser, DbPassword, DbHost, DbPort, DbName)
}