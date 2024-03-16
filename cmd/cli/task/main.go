package main

import (
	crand "crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"math/rand"

	"hiyoko-echo/configs"
	"hiyoko-echo/internal/infrastructure/database"
	"hiyoko-echo/pkg/logging/file"
	"hiyoko-echo/util"
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
		logger.Fatal("failed to create seed", "error", err)
	}
	rand.NewSource(seed.Int64())

	// flag
	server := flag.String("server", "local", "server environment")
	key = flag.String("key", "hiyoko", "exec task key")
	flag.Parse()

	// load env
	serverEnv = util.ServerEnv(*server)
	if ok := serverEnv.Regexp(); !ok {
		logger.Fatal("invalid server environment")
	}
	util.LoadEnv(serverEnv, EnvRoot)
	databaseConf = configs.NewMySqlConf()

	// timezone
	util.LoadTimezone()
}

func main() {
	entClient, err := database.NewMySqlConnect(serverEnv, databaseConf)
	if err != nil {
		logger.Fatal("failed to create dbclient", "error", err)
	}
	defer func(entClient *database.EntClient) {
		err := entClient.Close()
		if err != nil {
			logger.Fatal("failed to close dbclient", "error", err)
		}
	}(entClient)

	switch *key {
	case TaskKeyHiyoko:
		fmt.Println("hiyoko")
	}
}
