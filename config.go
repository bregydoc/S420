package s420

import (
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Config is the struct to config your s420
type Config struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	UseSSL          bool   `yaml:"use_ssl"`
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

	if strings.HasPrefix(c.Endpoint, "$") {
		c.Endpoint = os.Getenv(strings.Replace(c.Endpoint, "$", "", -1))
	}

	if strings.HasPrefix(c.AccessKeyID, "$") {
		c.AccessKeyID = os.Getenv(strings.Replace(c.AccessKeyID, "$", "", -1))
	}

	if strings.HasPrefix(c.SecretAccessKey, "$") {
		c.SecretAccessKey = os.Getenv(strings.Replace(c.SecretAccessKey, "$", "", -1))
	}

	return c, nil
}
