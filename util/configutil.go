package util

import (
	"strings"
)

type configUtil struct {
}

func (c *configUtil) Get(key string, defaultValue string) string {
	if key == "contextPath" {
		return "/api"
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
		return r
	}
	return path
}

var ConfigUtil configUtil = configUtil{}
