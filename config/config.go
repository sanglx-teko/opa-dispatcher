package config

import (
	"os"

	"encoding/json"

	etcdClient "github.com/coreos/etcd/clientv3"
)

var globalConfig *Configuration

// GetConfigurations get current config from system
func GetConfigurations() *Configuration {
	return globalConfig
}

// LoadConfigurations load configurations from json file
func LoadConfigurations(jsonFilePath string) error {
	// Try to open file config
	file, err := os.Open(jsonFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode json config
	decoder := json.NewDecoder(file)
	configure := Configuration{}
	if err := decoder.Decode(&configure); err != nil {
		return err
	}

	// Set global config
	globalConfig = &configure

	return nil
}

type (
	// Configuration ...
	Configuration struct {
		Service []Service
		ETCD    etcdClient.Config
	}

	// Service ...
	Service struct {
		Name         string
		ServiceGroup string
	}
)
