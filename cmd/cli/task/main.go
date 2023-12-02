package main

import (
	crand "crypto/rand"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"

	"hiyoko-echo/cmd/util"
	"hiyoko-echo/conf"
	"hiyoko-echo/infrastructure/database"
)

const (
	EnvRoot       = "cmd/cli"
	TaskKeyHiyoko = "hiyoko"
)

var (
	serverEnv    util.ServerEnv
	databaseConf database.Conf
	key          *string
)

func init() {
	// seed
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		log.Panicf("failed to create seed; error: %v", err)
	}
	rand.NewSource(seed.Int64())

	// flag
	server := flag.String("server", "local", "server environment")
	key = flag.String("key", "hiyoko", "exec task key")
	flag.Parse()

	// load env
	serverEnv = util.ServerEnv(*server)
	if ok := serverEnv.Regexp(); !ok {
		panic("invalid server environment")
	}
	util.LoadEnv(serverEnv, EnvRoot)
	databaseConf = conf.NewMySqlConf()

	// timezone
	util.LoadTimezone()
}

func main() {
	entClient, err := database.NewMySqlConnect(serverEnv, databaseConf)
	if err != nil {
		log.Panicf("failed to create dbclient; error: %v", err)
	}
	defer func(entClient *database.EntClient) {
		err := entClient.Close()
		if err != nil {
			log.Panicf("failed to close dbclient; error: %v", err)
		}
	}(entClient)

	switch *key {
	case TaskKeyHiyoko:
		fmt.Println("hiyoko")
	}
}
