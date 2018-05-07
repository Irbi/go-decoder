// Provides connection handler for Redis database
// Implements manager.DataStore interface
package redis

import (
	"github.com/go-redis/redis"
	"fmt"
	"strconv"
)

// Describes connection configuration
// Contains connection handler after initialization
type Connector struct {
	Host string
	Port string
	User string
	Password string
	Dbname string
	DB *redis.Client
}

// Constructor that accept a common argument
func NewConnector(conf map[string]string) *Connector {
	c := Connector{}

	c.SetHost(conf["Host"])
	c.SetPort(conf["Port"])
	c.SetUser(conf["User"])
	c.SetPassword(conf["Password"])
	c.SetDbname(conf["Dbname"])

	c.Connect()

	return &c
}

func (c *Connector) SetHost(host string) {
	c.Host = host
}

func (c *Connector) SetPort(port string) {
	c.Port = port
}

func (c *Connector) SetUser(user string) {
	c.User = user
}

func (c *Connector) SetPassword(pwd string) {
	c.Password = pwd
}

func (c *Connector) SetDbname(dbname string) {
	c.Dbname = dbname
}

// Connects to database and fills DB field with handler
func (c *Connector) Connect() {
	dbName, _ := strconv.Atoi(c.Dbname)
	port, _ := strconv.Atoi(c.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, port),
		Password: c.Password,
		DB:       dbName,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}

	c.DB = client
}

func (c *Connector) Name() string {
	return "RedisDataStore"
}