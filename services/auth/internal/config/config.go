package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppName      string
	AppEnv       string
	GRPCPort     string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	JWTSecret    string
	KafkaBrokers string
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) && !os.IsNotExist(err) {
			return nil, err
		}
	}
	return &Config{
		AppName:      viper.GetString("APP_NAME"),
		AppEnv:       viper.GetString("APP_ENV"),
		GRPCPort:     viper.GetString("GRPC_PORT"),
		DBHost:       viper.GetString("DB_HOST"),
		DBPort:       viper.GetString("DB_PORT"),
		DBUser:       viper.GetString("DB_USER"),
		DBPassword:   viper.GetString("DB_PASSWORD"),
		DBName:       viper.GetString("DB_NAME"),
		DBSSLMode:    viper.GetString("DB_SSLMODE"),
		JWTSecret:    viper.GetString("JWT_SECRET"),
		KafkaBrokers: viper.GetString("KAFKA_BROKERS"),
	}, nil
}
