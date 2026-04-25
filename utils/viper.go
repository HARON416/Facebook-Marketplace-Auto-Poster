package utils

import (
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	GEMINI_API_KEY string `mapstructure:"GEMINI_API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	v := viper.New()
	v.SetConfigType("env")
	v.SetConfigFile(filepath.Join(path, ".env"))

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	err = v.Unmarshal(&config)

	return
}
