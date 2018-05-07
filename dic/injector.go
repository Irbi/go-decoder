// DependencyContainer build and stores the application-wide dependencies
package dic

import (
	"decoder/app"
	"fmt"
	"decoder/manager"
)

// DependencyContainer stores the application-wide dependencies
// ReadStorage provides handler for the application read storage -- redis, mysql etc
// WriteStorage provides handler for the application write storage -- redis, mysql etc
// Daos contains specific DAO instances basing on ReadStorage and WriteStorage settings
// Logger contains application Logger
type DependencyContainer struct {
	ReadStorage 	manager.DataStore
	WriteStorage 	manager.DataStore
	Daos			map[interface{}]interface{}
	Logger 			app.Logger
}

// Returns dependency container with initialized services
func Load() DependencyContainer {
	di := DependencyContainer{}

	di.setReadStorage()
	di.setWriteStorage()

	// Other dependencies here

	return di
}

// Map read storage dependencies
func (dc *DependencyContainer) setReadStorage () {

	storage := app.AppConfig.Decoder.ReadStorage.Engine
	handler := buildDataStorage(storage, "read")

	dc.ReadStorage = handler
}

// Map write storage dependencies
func (dc *DependencyContainer) setWriteStorage () {
	storage := app.AppConfig.Decoder.WriteStorage.Engine
	handler := buildDataStorage(storage, "write")

	dc.ReadStorage = handler
}

// Private method, encapsulates calls to data storages factory
func buildDataStorage(storage string, stype string) manager.DataStore {
	stCfg := app.StorageConfig.Database[storage]
	config:= stCfg.ConvertToMap()

	manager.Register(storage)

	datastore, err := manager.CreateDataStore(storage, config)
	if (err != nil) {
		panic(fmt.Errorf("Failed to load %s data storage: %s", stype, err)) // as usual, don't panic in your app
	}

	return datastore
}

// Map read DAOs dependencies providing DC ReadStorage as argument
func (dc *DependencyContainer) setReadDaos () {}

// Map write DAOs dependencies providing DC WriteStorage as argument
func (dc *DependencyContainer) setWriteDaos () {}

// Map application logger
func (dc *DependencyContainer) setLogger (l app.Logger) {}