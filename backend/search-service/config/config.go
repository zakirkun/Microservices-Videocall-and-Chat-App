package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config holds user service configuration
type Config struct {
	DBAddress  string `mapstructure:"dbAddress"`
	DBUser     string `mapstructure:"dbUser"`
	DBPassword string `mapstructure:"dbPassword"`

	HTTPPort string `mapstructure:"httpPort"`

	Origin string `mapstructure:"origin"`

	BrokerAddress string `mapstructure:"brokerAddress"`
}

// LoadConfigFromEnvironment loads user service configuration from environment variables and returns an error
// if any of them is missing
func LoadConfigFromEnvironment() (conf Config, err error) {

	conf.DBAddress = os.Getenv("ES_ADDRESS")
	if len(conf.DBAddress) == 0 {
		return Config{}, errors.New("Environment variable ES_ADDRESS not set")
	}
	conf.DBUser = os.Getenv("ES_USER")
	conf.DBPassword = os.Getenv("ES_PASS")

	conf.HTTPPort = os.Getenv("HTTP_PORT")
	if len(conf.HTTPPort) == 0 {
		return Config{}, errors.New("Environment variable HTTP_PORT not set")
	}

	conf.Origin = os.Getenv("ORIGIN")
	if len(conf.Origin) == 0 {
		return Config{}, errors.New("Environment variable ORIGIN not set")
	}

	conf.BrokerAddress = os.Getenv("BROKER_ADDRESS")
	if len(conf.BrokerAddress) == 0 {
		return Config{}, errors.New("Environment variable BROKER_ADDRESS not set")
	}

	return
}

// LoadConfigFromFile loads config from specified path
func LoadConfigFromFile(path string) (config Config, err error) {
	vp := viper.New()

	vp.AddConfigPath(filepath.Dir(path))

	filename := strings.Split(filepath.Base(path), ".")
	vp.SetConfigName(filename[0])
	vp.SetConfigType(filename[1])

	if err = vp.ReadInConfig(); err != nil {
		return
	}

	err = vp.Unmarshal(&config)
	return
}
