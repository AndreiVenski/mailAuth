package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	Postgres PostgresqlConfig
	Server   ServerConfig
	Test     TestConfig
}

type TestConfig struct {
	AddTestDataToDB      bool   `envconfig:"TEST_ADDTESTDATATODB"`
	TestDataScriptSource string `envconfig:"TEST_TESTDATASCRIPTSOURCE"`
}

type PostgresqlConfig struct {
	PostgresqlHost     string `envconfig:"POSTGRESQL_HOST"`
	PostgresqlPort     string `envconfig:"POSTGRESQL_PORT"`
	PostgresqlUser     string `envconfig:"POSTGRESQL_USER"`
	PostgresqlDbname   string `envconfig:"POSTGRESQL_DBNAME"`
	PostgresqlPassword string `envconfig:"POSTGRESQL_PASSWORD"`
}

type ServerConfig struct {
	RunningPort         string        `envconfig:"SERVER_RUNNINGPORT"`
	JWTSecret           string        `envconfig:"SERVER_JWTSECRET"`
	AccessTokenExpires  time.Duration `envconfig:"SERVER_ACCESSTOKENEXPIRES"`
	RefreshTokenExpires time.Duration `envconfig:"SERVER_REFRESHTOKENEXPIRES"`
}

func InitConfig(path string) (*Config, error) {
	var cfg Config
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	if err = envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
