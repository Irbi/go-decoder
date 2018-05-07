// Provides connection handler for MySQL storage
// Implements manager.DataStore interface
package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

// Describes connection configuration
// Contains connection handler after initialization
type Connector struct {
	Host string
	Port string
	User string
	Password string
	Dbname string
	DB *sql.DB
}

// Constructor that accept a common argument
func NewConnector(conf map[string]string) *Connector {
	// you can use something like this:
	//c := Connector{
	// Host:conf["Host"],
	// Port:conf["Port"],
	// etc...
	// }
	// But maybe you need tests or validation before connection, so:
	c := Connector{}

	c.SetHost(conf["Host"])
	c.SetPort(conf["Port"])
	c.SetUser(conf["User"])
	c.SetPassword(conf["Password"])
	c.SetDbname(conf["Dbname"])

	c.Connect()

	return &c
}

// Add tests and validations
func (c *Connector) SetHost(host string) {
	c.Host = host
}

// Add tests and validations
func (c *Connector) SetPort(port string) {
	c.Port = port
}

// Add tests and validations
func (c *Connector) SetUser(user string) {
	c.User = user
}

// Add tests and validations
func (c *Connector) SetPassword(pwd string) {
	c.Password = pwd
}

// Add tests and validations
func (c *Connector) SetDbname(dbname string) {
	c.Dbname = dbname
}

// Connects to database and fills DB field with handler
func (c *Connector) Connect() {
	connString := fmt.Sprintf("%s:%s@/%s", c.User, c.Password, c.Dbname)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	c.DB = db
}

// Just example
func (c *Connector) Name() string {
	return "MySQLDataStore"
}