package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Config can be used to get a setting
type Config struct {
	Settings map[string]interface{}
}

// NewConfig reads a config file and returns a receiver to get keys.
func NewConfig(filename string) *Config {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Panicf("Unable to open config file: %+v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	c := &Config{
		Settings: result,
	}

	return c
}

// Get will look up a key and, if it exists, will return it
func (c *Config) Get(key string) interface{} {
	parts := strings.Split(key, ".")

	if len(parts) > 0 {
		var s map[string]interface{}
		for i := 0; i < len(parts)-1; i++ {
			s = c.Settings[parts[i]].(map[string]interface{})
		}

		return s[parts[len(parts)-1]]
	}

	return c.Settings[key]
}
