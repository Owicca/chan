package infra

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConn(DbHost string, DbPort string, DbName string, DbUser string, DbPassword string) (*gorm.DB, error) {
	connectionString := GetConnString("mysql", DbHost, DbPort, DbName, DbUser, DbPassword)
	dialector := GetDialector("mysql", connectionString)
	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Database connection failed %s", err)
	}

	var tableList []string
	conn.Raw("SHOW TABLES LIKE 'posts'").Scan(&tableList)
	if len(tableList) == 0 {
		PopulateDb(conn)
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

func PopulateDb(db *gorm.DB) {
	data, _ := os.ReadFile("./db_schema.my.sql")

	db.Exec(string(data))
}

func ClearDb(db *gorm.DB) {
	tables := []string{
		"pair_to_role",
		"action_to_object",

		"posts",
		"threads",
		"boards",
		"topics",
		"users",

		"roles",
		"log_actions",
		"objects",
		"actions",
		"media",
	}

	for _, name := range tables {
		db.Exec("DELETE FROM " + name)
	}
}
