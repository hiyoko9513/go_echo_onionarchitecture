package util

import (
	"fmt"
	"github.com/joho/godotenv"
	"hiyoko-echo/shared"
	"log"
	"regexp"
	"time"
)

type ServerEnv string

func (s *ServerEnv) String() string {
	return string(*s)
}

func (s *ServerEnv) Regexp() bool {
	regex := regexp.MustCompile("^local$|^develop$|^staging$|^production$")
	match := regex.MatchString(s.String())
	if !match {
		log.Fatal("invalid server environment")
	}
	return true
}

func LoadEnv(server ServerEnv, rootPath string) {
	envPath := fmt.Sprintf("%s/.env.%s", rootPath, server)
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadTimezone() {
	time.Local = shared.Timezone()
}
