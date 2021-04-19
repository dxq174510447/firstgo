package util

import (
	"fmt"
	"os"
	"strings"
)

type configUtil struct {
}

func (c *configUtil) Get(key string, defaultValue string) string {
	if key == "contextPath" {
		return "/api"
	}
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return defaultValue
}

func (c *configUtil) ClearHttpPath(path string) string {
	return path
}

func (c *configUtil) RemovePrefix(path string, prefix string) string {

	//	fmt.Println(path,prefix)
	if path == prefix {
		return ""
	}

	if strings.HasPrefix(path, prefix) {
		r := path[len(prefix):]
		if r == "/" {
			return ""
		}
		if r[0:1] != "/" {
			return fmt.Sprintf("/%s", r)
		}
		return r
	}
	return path
}

var ConfigUtil configUtil = configUtil{}
