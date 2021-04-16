package util

import "strings"

type configUtil struct {
}

func (c *configUtil) Get(key string, defaultValue string) string {
	if key == "contextPath" {
		return "/api"
	}
	return ""
}

func (c *configUtil) ClearHttpPath(path string) string {
	return path
}

func (c *configUtil) RemovePrefix(path string, prefix string) string {
	if path == prefix {
		return ""
	}
	if strings.HasPrefix(path, prefix) {
		return path[len(prefix):]
	}
	return path
}

var ConfigUtil configUtil = configUtil{}
