package common

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MySqlDSN string `envconfig:"MYSQL_DSN"`

	GetAccountByIDQuery string `envconfig:"GET_ACCOUNT_BY_ID_QUERY"`
}

func LoadConfig() (Config, error) {
	config := Config{}
	err := envconfig.Process("", &config)
	return config, err
}