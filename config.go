package s420

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config is the struct to config your s420
type Config struct {
	Storage Storage `yaml:"storage"`
	Service Service `yaml:"service"`
	Public  Public  `yaml:"public"`
}

// Storage define the config to storage system (default: minio)
type Storage struct {
	Backend         string `yaml:"backend"`
	Endpoint        string `yaml:"endpoint"`
	AccessKey       string `yaml:"access_key"`
	SecretAccessKey string `yaml:"secret_access_key"`
	UseSSL          bool   `yaml:"use_ssl"`
}

// Service define the grpc service configuration
type Service struct {
	Port int64 `yaml:"port"`
}

// Public is the public configuration
type Public struct {
	Prefix string `yaml:"prefix"`
	Port   int64  `yaml:"port"`
}

// NewConfigFromFile returns a new config struct
func NewConfigFromFile(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	c := new(Config)
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
