// The main entry of the application
//
// models: contains the data structures used for communication between different layers
// services: the main business logic of the application
// redis: contains the DAO (Data Access Object) layer that interacts with the Redis data storage
// mysql: contains the DAO (Data Access Object) layer that interacts with the MySQL data storage
// handlers: contains the API layer that wires up the HTTP routes with the corresponding service APIs
// manager: the building yard for data storage connection handlers and DAOs
// api: third-party API clients
//
// Used packages:
// routing -- https://github.com/go-ozzo/ozzo-routing
// config parser and formatter -- https://github.com/BurntSushi/toml
// redis -- https://github.com/go-redis/redis
// mysql -- https://github.com/go-sql-driver/mysql

package main

import (
	"fmt"
	"net/http"
	"decoder/app"
	"decoder/dic"
	"github.com/go-ozzo/ozzo-routing"
	"os"
)

const configsPath string = "src/decoder/configs"

// Starting point of application. Responsibilities:
//
// - load configuration
// - initialize Dependency Container basing on configuration file(s)
// - instantiate components and inject dependencies
// - initialize third-party apis
// - start the HTTP server
func main() {
	app.LoadConfig(configsPath)
	coreDeps := dic.Load()

	fmt.Println(coreDeps.ReadStorage)
	os.Exit(0)

	// wire application routing
	http.Handle("/", buildRouter(coreDeps))

	// start the server
	address := fmt.Sprintf("%v", app.AppConfig.Server.Address)
	panic(http.ListenAndServe(address, nil))
}

// Init routing using dependencies
func buildRouter(deps dic.DependencyContainer) *routing.Router {
	router := routing.New()

	//rg := router.Group("/v1")

	// This added for only purpose to demonstrate router/router groups using
	//vinDAO := redis.NewVinDAO()
	//handlers.ServeDecoderResource(rg, services.NewDecoderService(vinDAO))

	// add more resource APIs if needed

	return router
}