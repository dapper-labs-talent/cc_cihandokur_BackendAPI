package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Postgres *Postgres `validate:"required"`
	Server   *Server
	Log      *Log `validate:"required"`
	Jwt      *Jwt
}

type Jwt struct {
	SecretKey  string `validate:"required"`
	ExpDurHour uint16 `validate:"required"`
	Header     string `validate:"required"`
}

type Postgres struct {
	Database      string `validate:"required"`
	Host          string `validate:"required"`
	Port          string `validate:"required"`
	Pass          string `validate:"required"`
	User          string `validate:"required"`
	LocalDataPath string
}

type Server struct {
	Host string
	Port uint `validate:"max=65535"`
}

type Log struct {
	Level string `validate:"oneof=debug info warn error fatal"`
}

var Config *Configuration

func LoadConfiguration() error {

	viper.SetConfigType("toml")

	// file name
	viper.SetConfigName("config")

	// paths
	viper.AddConfigPath("/config/")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./app/config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("/go/bin/devlab")

	viper.AllowEmptyEnv(true)

	// read config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	// unmarshal config
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("error unmarshall config: %v", err)
	}

	// JWT
	if err := viper.BindEnv("jwt.key", "JWT_SECRET_KEY"); err != nil {
		log.Fatalf("error binding env var `JWT_SECRET_KEY`: %v", err)
	}
	if err := viper.BindEnv("jwt.expDurMin", "JWT_EXP_DUR_MIN"); err != nil {
		log.Fatalf("error binding env var `JWT_EXP_DUR_MIN`: %v", err)
	}

	// postgres
	if err := viper.BindEnv("postgres.db", "POSTGRES_DB"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_DB`: %v", err)
	}
	if err := viper.BindEnv("postgres.host", "POSTGRES_HOST"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_HOST`: %v", err)
	}
	if err := viper.BindEnv("postgres.pass", "POSTGRES_PASSWORD"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_PASSWORD`: %v", err)
	}
	if err := viper.BindEnv("postgres.port", "POSTGRES_PORT"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_PORT`: %v", err)
	}
	if err := viper.BindEnv("postgres.user", "POSTGRES_USER"); err != nil {
		log.Fatalf("error binding env var `POSTGRES_USER`: %v", err)
	}

	// server
	if err := viper.BindEnv("server.host", "SERVER_HOST"); err != nil {
		log.Fatalf("error binding env var `SERVER_HOST`: %v", err)
	}
	if err := viper.BindEnv("server.port", "SERVER_PORT"); err != nil {
		log.Fatalf("error binding env var `SERVER_PORT`: %v", err)
	}

	// log
	if err := viper.BindEnv("log.level", "LOG_LEVEL"); err != nil {
		log.Fatalf("error binding env var `LOG_LEVEL`: %v", err)
	}

	return nil
}
