package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App     App
	MysqlDB MysqlDB
	MongoDB MongoDB
}

type App struct {
	Name     string `json:"name"`
	HttpPort string `json:"http_port"`
}

type MysqlDB struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

type MongoDB struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func LoadConfiguration() Config {
	cfg := Config{
		App: App{
			Name:     viper.GetString("app.name"),
			HttpPort: viper.GetString("app.port"),
		},
		MysqlDB: MysqlDB{
			Name:     viper.GetString("mysql.name"),
			Host:     viper.GetString("mysql.host"),
			User:     viper.GetString("mysql.user"),
			Port:     viper.GetString("mysql.port"),
			Password: viper.GetString("mysql.password"),
		},
		MongoDB: MongoDB{
			Name:     viper.GetString("mongo.name"),
			Host:     viper.GetString("mongo.host"),
			User:     viper.GetString("mongo.user"),
			Port:     viper.GetString("mongo.port"),
			Password: viper.GetString("mongo.password"),
		},
	}

	return cfg
}
