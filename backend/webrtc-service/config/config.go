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
	PodName      string `mapstructure:"podName"`
	PodNamespace string `mapstructure:"podNamespace"`
	ServiceName  string `mapstructure:"serviceName"`

	TURNAddress  string `mapstructure:"turnAddress"`
	TURNPort     string `mapstructure:"turnPort"`
	TURNSPort    string `mapstructure:"turnsPort"`
	TURNUser     string `mapstructure:"turnUser"`
	TURNPassword string `mapstructure:"turnPassword"`

	DBAddress  string `mapstructure:"dbAddress"`
	DBPassword string `mapstructure:"dbPassword"`

	HTTPPort string `mapstructure:"httpPort"`

	Origin string `mapstructure:"origin"`

	BrokerAddress string `mapstructure:"brokerAddress"`
}

// LoadConfigFromEnvironment loads user service configuration from environment variables and returns an error
// if any of them is missing
func LoadConfigFromEnvironment() (conf Config, err error) {

	conf.PodName = os.Getenv("POD_NAME")
	if len(conf.PodName) == 0 {
		return Config{}, errors.New("Environment variable POD_NAME not set")
	}
	conf.PodNamespace = os.Getenv("POD_NAMESPACE")
	if len(conf.PodNamespace) == 0 {
		return Config{}, errors.New("Environment variable POD_NAMESPACE not set")
	}
	conf.ServiceName = os.Getenv("SERVICE_NAME")
	if len(conf.ServiceName) == 0 {
		return Config{}, errors.New("Environment variable SERVICE_NAME not set")
	}

	conf.TURNAddress = os.Getenv("TURN_ADDRESS")
	if len(conf.TURNAddress) == 0 {
		return Config{}, errors.New("Environment variable TURN_ADDRESS not set")
	}
	conf.TURNPort = os.Getenv("TURN_PORT")
	if len(conf.TURNPort) == 0 {
		return Config{}, errors.New("Environment variable TURN_PORT not set")
	}
	conf.TURNSPort = os.Getenv("TURN_TLS_PORT")
	if len(conf.TURNSPort) == 0 {
		return Config{}, errors.New("Environment variable TURN_TLS_PORT not set")
	}
	conf.TURNUser = os.Getenv("TURN_USER")
	if len(conf.TURNUser) == 0 {
		return Config{}, errors.New("Environment variable TURN_USER not set")
	}
	conf.TURNPassword = os.Getenv("TURN_PASSWORD")
	if len(conf.TURNPassword) == 0 {
		return Config{}, errors.New("Environment variable TURN_PASSWORD not set")
	}

	conf.DBAddress = os.Getenv("REDIS_ADDRESS")
	if len(conf.DBAddress) == 0 {
		return Config{}, errors.New("Environment variable REDIS_ADDRESS not set")
	}
	conf.DBPassword = os.Getenv("REDIS_PASSWORD")
	if len(conf.DBPassword) == 0 {
		return Config{}, errors.New("Environment variable REDIS_PASSWORD not set")
	}

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
