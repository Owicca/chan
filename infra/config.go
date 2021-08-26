package infra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Config struct {
	HttpHost   string
	HttpPort   string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	Logger     zap.Config
	StoreKey   string
}

func LoadConfig(path string) (Config, error) {
	cfg := NewConfig()

	file, err := os.Open(path)
	if err != nil {
		return cfg, fmt.Errorf("Could not open config file %s (%s)", path, err)
	}
	dec := json.NewDecoder(file)
	if err := dec.Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("Could not decode contents of config file %s (%s)", path, err)
	}

	return cfg, nil
}

func NewConfig() Config {
	var cfg Config

	return cfg
}

func Setup(configPath string) (Config, *pgxpool.Pool, *zap.Logger) {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("{\"level\":\"fatal\", \"message\":\"Could not load configuration (%s)\", \"timestamp\": %d", err, time.Now().Unix())
	}
	logger, err := cfg.Logger.Build()
	if err != nil {
		log.Fatalf("{\"level\":\"fatal\", \"message\":\"Can't initialize zap logger (%s)\", \"timestamp\": %d", err, time.Now().Unix())
	}

	conn, err := GetDbConn(cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbUser, cfg.DbPassword)
	if err != nil {
		errMsg := fmt.Sprintf("Error while connecting to db (%s)", err)
		logger.Fatal(errMsg, zap.Int64("timestamp", time.Now().Unix()))
	}

	return cfg, conn, logger
}