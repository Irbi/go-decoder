// DAOs factory.
// Creates DAO objects depending on given congiguration (AppConfig.ReadStorage, AppConfig.WriteStorage,)
// All additional storage interfaces must be implemented in the similar way
// Implement here DAOs fabric using "manager.storage.go" as an example
package manager

// DaoStore defines the list of methods that must be implemented in each DAO.
// Method "Name" just returns the name of the instance
type DaoStore interface {
	Name() string
}

// Common interface for connection handlers constructors
type DaoStoreFactory func(conf map[string]string) (DaoStore, error)

// Map of the registered factories
var daoStoreFactories = make(map[string]DataStoreFactory)
