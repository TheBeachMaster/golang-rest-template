package config

import (
	"errors"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	DB     PostgresConfig
	Cache  RedisConfig
	Server ServerConfig
	Logger LoggerConfig
	Port   string `envconfig:"PORT"`
}

type RedisConfig struct {
	RedisURL string `envconfig:"REDIS_URL"`
}

type PostgresConfig struct {
	DataBaseURL string `envconfig:"DATABASE_URL"`
}

type ServerConfig struct {
	ReadTimeout  uint8  `envconfig:"READTIMEOUT"`
	AppName      string `envconfig:"APPNAME"`
	ServerHeader string `envconfig:"SERVERHEADER"`
	Prefork      bool   `envconfig:"PREFORK"`
}

type LoggerConfig struct {
	LogFile           string `envconfig:"LOGFILE"`
	Timezone          string `envconfig:"TIMEZONE"`
	LoggerMaxFileSize uint8  `envconfig:"LOGGER_MAX_FILE_SIZE"`
	CompressLogFile   bool   `envconfig:"COMPRESS_LOG_FILE"`
	UseLocalTime      bool   `envconfig:"USE_LOCALTIME"`
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

func LoadEnvConfig() (*Config, error) {
	var conf Config
	configErr := envconfig.Process("", &conf)
	if configErr != nil {
		log.Printf("Unable to load env config due to %s", configErr.Error())
		return nil, configErr
	}
	return &conf, nil
}
