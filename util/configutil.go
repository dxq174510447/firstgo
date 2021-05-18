package util

import (
	"fmt"
	"os"
)

type configUtil struct {
}

// WrapServletPath 暂时没想到更好的办法 提前把servletPath 设置到dispatcher中，临时解决
func (c *configUtil) WrapServletPath(path string) string {
	sp := c.Get("contextPath", "/api")
	return fmt.Sprintf("%s%s", sp, path)
}

func (c *configUtil) Get(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return defaultValue
}

var ConfigUtil configUtil = configUtil{}
