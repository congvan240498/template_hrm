package config

import (
	"github.com/caarlos0/env/v7"
	"hrm/pkg/logger"
)

type SystemConfig struct {
	Env      string `env:"ENV,required,notEmpty"`
	HttpPort uint64 `env:"HTTP_PORT,required,notEmpty"`

	MongoDBConfig MongoDBConfig `envPrefix:"MONGODB_"`
	SecretKey     string        `env:"SECRET_KEY,required,notEmpty"`
}

type MongoDBConfig struct {
	DatabaseURI  string `env:"DATABASE_URI,required,notEmpty"`
	DatabaseName string `env:"DATABASE_NAME,required,notEmpty"`
}

var configSingletonObj *SystemConfig

func LoadConfig() (cf *SystemConfig, err error) {
	log := logger.GetLogger()

	if configSingletonObj != nil {
		cf = configSingletonObj
		return
	}

	cf = &SystemConfig{}
	if err = env.Parse(cf); err != nil {
		log.Fatal().Err(err).Msg("failed to unmarshal config")
	}

	configSingletonObj = cf
	return
}

func GetInstance() *SystemConfig {
	return configSingletonObj
}
