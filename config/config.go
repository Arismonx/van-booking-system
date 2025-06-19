package config

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   *Server   `mapstructure:"server" validate:"required"`
		Database *Database `mapstructure:"database" validate:"required"`
	}

	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     string `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}

	Server struct {
		Port           int           `mapstructure:"port" validate:"required"`
		AllowedOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit      string        `mapstructure:"bodyLimit" validate:"required"`
		TimeOut        time.Duration `mapstructure:"timeout" validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func LoadConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic("Error read in config :" + err.Error())
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic("Error unmarshal mapstructure:" + err.Error())
		}

		validating := validator.New()
		if err := validating.Struct(configInstance); err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				log.Fatal(err)
			}
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println("Field:", err.Field())
				fmt.Println("Tag:", err.Tag())
				fmt.Println("Value:", err.Value())
				fmt.Println("Param:", err.Param())
				fmt.Println("---")
			}
			panic("Error invalid!:" + err.Error())
		} else {
			fmt.Println("valid!")
		}
	})

	return configInstance
}
