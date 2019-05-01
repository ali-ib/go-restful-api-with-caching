package configs

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

// Go represntation of configs
type Config struct {
	CONNECTIONSTRING string
	DBNAME           string
	COLLECTIONNAME   string
	CACHETIME        time.Duration
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

var Configs Config

func init() {
	Configs.Read()
}
