package app

import (
	"github.com/BurntSushi/toml"
	"fmt"
	"reflect"
)

// AppConfig stores the application-wide configurations: current strategy, read/write storages etc
var AppConfig AppConfigStruct
// StorageConfig stores the data storages settings. Only storages defined in AppConfig will be used
var StorageConfig StorageConfigStruct

type AppConfigStruct struct {
	Server struct {
		Address string
	}
	Auth struct {
		User, Password string
	}
	Decoder struct {
		StrategyPrefix string
		ReadStorage struct {
			Type string
			Engine string
		}
		WriteStorage struct {
			Type string
			Engine string
		}
	}
}

type StorageConfigStruct struct {
	Database map[string]DbStorageEngine
}

type DbStorageEngine struct {
	Host string
	Port string
	User string
	Password string
	Dbname string
	Handler interface{}
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be .toml
func LoadConfig(configsPath string) {

	if _, err := toml.DecodeFile(configsPath + "/app.toml", &AppConfig); err != nil {
		panic(fmt.Errorf("Failed to read the configuration file: %s", err))
	}
	if _, err := toml.DecodeFile(configsPath + "/storage.toml", &StorageConfig); err != nil {
		panic(fmt.Errorf("Failed to read the configuration file: %s", err))
	}

	AppConfig.Validate()
	StorageConfig.Validate()
}

// Add validation logic for AppConfig if needed
func (config AppConfigStruct) Validate() error {
	return nil
}

// Add validation logic for StorageConfig if needed
func (config StorageConfigStruct) Validate() error {
	return nil
}

func (config DbStorageEngine) ConvertToMap() map[string]string {
	r := reflect.ValueOf(config)
	m := make(map[string]string, r.NumField())

	for i := 0; i < r.NumField(); i++ {
		m[r.Type().Field(i).Name] = r.Field(i).Interface().(string)
	}

	return m
}


