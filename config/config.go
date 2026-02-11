package config

import (
	"os"

	"github.com/gookit/slog"
	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	Pprof  Pprof
}

type Server struct {
	Host string
	Port int
}

type Pprof struct {
	Host string
	Port int
}

func GetConfig() Config {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		slog.Fatal("Failed to read .env file", "error", err)
		os.Exit(1)
	}

	return Config{
		Server: Server{
			Host: viper.GetString("SRV_HOST"),
			Port: viper.GetInt("SRV_PORT"),
		},
		Pprof: Pprof{
			Host: viper.GetString("PPROF_HOST"),
			Port: viper.GetInt("PPROF_PORT"),
		},
	}
}
