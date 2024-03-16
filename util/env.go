package util

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"strconv"
)

type ServerEnv string

func (s *ServerEnv) String() string {
	return string(*s)
}

func (s *ServerEnv) Regexp() bool {
	regex := regexp.MustCompile("^local$|^develop$|^staging$|^production$")
	match := regex.MatchString(s.String())
	if !match {
		log.Panicf("not valid environment; env: %s", s.String())
	}
	return true
}

func LoadEnv(server ServerEnv, rootPath string) {
	envPath := fmt.Sprintf("%s/.env.%s", rootPath, server)
	err := godotenv.Load(envPath)
	if err != nil {
		log.Panicf("failed to load environment; error: %v", err)
	}
}

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
