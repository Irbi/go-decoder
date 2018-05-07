// Data storages factory.
// Creates storage objects without having to specify the exact type
// All additional storage interfaces must be implemented in the similar way
package manager

import (
	"decoder/redis"
	"log"
	"errors"
	"fmt"
	"strings"
	"decoder/mysql"
)

// DataStore defines the list of methods that must be implemented in each data storage.
// Add here additional methods (ex TestConnection) if needed and don't forget to
// implement them in the storages packages
// Method "Name" just returns the name of the instance
type DataStore interface {
	Name() string
}

// Common interface for connection handlers constructors
type DataStoreFactory func(conf map[string]string) (DataStore, error)

// Map of the registered factories
var datastoreFactories = make(map[string]DataStoreFactory)

// Maps factories to constructors
var datastoreConstructors = map[string]DataStoreFactory{
	"mysql": NewMySQLDataStore,
	"redis": NewRedisDataStore,
}

// Returns connection handler for Redis database using the given configuration map
// Uses type redis.Connector (package redis).
// All methods defined in DataStore interface are implemented in mysql package
func NewRedisDataStore(conf map[string]string) (DataStore, error) {
	connector := redis.NewConnector(conf)
	return connector, nil
}

// Returns connection handler for MySQL database using the given configuration map
// Uses type mysql.Connector (package mysql).
// All methods defined in DataStore interface are implemented in redis package
func NewMySQLDataStore(conf map[string]string) (DataStore, error) {
	connector := mysql.NewConnector(conf)
	return connector, nil
}

// Helper method to add factories to the common map "datastoreFactories"
func RegisterDataStore(storageName string, factory DataStoreFactory) {
	if factory == nil {
		log.Panicf("Datastore factory %s does not exist.", storageName)
	}
	_, registered := datastoreFactories[storageName]
	if registered {
		log.Printf("Datastore factory %s already registered. Ignoring.", storageName)
	}
	datastoreFactories[storageName] = factory
}

// Maps registered factories to constructors and registers
func Register(storageName string) {
	RegisterDataStore(storageName, datastoreConstructors[storageName])
}

// CreateDataStore calls the appropriate factory method using
// the conf argument "storageName" to define instance
// and conf argument "conf" to create an instance of the DataStore interface.
func CreateDataStore(storageName string, conf map[string]string) (DataStore, error) {
	engineFactory, ok := datastoreFactories[storageName]
	if !ok {
		// Factory has not been registered
		// Make a list of all available datastore factories for logging
		availableDatastores := make([]string, len(datastoreFactories))
		for k, _ := range datastoreFactories {
			availableDatastores = append(availableDatastores, k)
		}
		return nil, errors.New(fmt.Sprintf("Invalid Datastore name. Must be one of: %s", strings.Join(availableDatastores, ", ")))
	}

	// Run the factory with the configuration.
	return engineFactory(conf)
}
