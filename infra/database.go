package infra

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetDbConn(DbHost string, DbPort string, DbName string, DbUser string, DbPassword string) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		DbHost, DbPort, DbName, DbUser, DbPassword)
	conn, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("Database connection failed %s", err)
	}

	return conn, nil
}
