package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Postgres PostgresConfig
	Redis    RedisConfig
	Env      string
	HttpAddr string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func MustLoad() *Config {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: %s", err)
	}
	viper.AddConfigPath(os.Getenv("CONFIG_PATH"))
	viper.SetConfigName(os.Getenv("CONFIG_NAME"))
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Ошибка при загрузке конфига: %s", err)
	}
	return newConfig()
}

func newConfig() *Config {
	redisDb, _ := strconv.Atoi(viper.GetString("redis.db"))
	return &Config{
		Postgres: PostgresConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: viper.GetString("db.password"),
		},
		Redis: RedisConfig{
			Addr:     viper.GetString("redis.addr"),
			Password: "",
			DB:       redisDb,
		},
		Env:      viper.GetString("env"),
		HttpAddr: fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")),
	}
}
