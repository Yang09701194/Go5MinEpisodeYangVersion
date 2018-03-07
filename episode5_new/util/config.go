package util

import "fmt"

type Env string

func (e Env) String() string {
	return string(e)
}

const (
	EnvDev   Env = "dev"
	EnvStage Env = "stage"
	EnvProd  Env = "prod"
)

type Config struct {
	Port        int    `envconfig:"port"`
	Environment string `envconfig:"envi"`
	MongoURL    string `envconfig:mongo_url`
}

func (c Config) Env() (Env, error) {
	switch c.Environment {
	//dev or stage or prod
	//use Environment of Config through case turn into Env
	case EnvDev.String():
		return EnvDev, nil
	case EnvStage.String():
		return EnvDev, nil
	case EnvProd.String():
		return EnvProd, nil
	default:
		return Env(""), fmt.Errorf("invalid environment %s", c.Environment)
	}
}
