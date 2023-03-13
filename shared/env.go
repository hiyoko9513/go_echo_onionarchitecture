package shared

import (
	"os"
	"strconv"
)

type Env string

func (e Env) GetString(defaultVal string) string {
	value := os.Getenv(string(e))
	if value == "" {
		return defaultVal
	}
	return value
}

func (e Env) GetInt(defaultVal int) int {
	valString := e.GetString("")
	if valString == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valString)
	if err != nil {
		return defaultVal
	}
	return val
}
