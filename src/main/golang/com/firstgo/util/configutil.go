package util

import (
	"os"
)

type configUtil struct {
}

func (c *configUtil) Get(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return defaultValue
}

var ConfigUtil configUtil = configUtil{}
