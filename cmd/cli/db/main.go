package main

import (
	"context"
	crand "crypto/rand"
	"flag"
	"fmt"
	"hiyoko-echo/util"
	"log"
	"math"
	"math/big"
	"math/rand"

	"hiyoko-echo/configs"
	"hiyoko-echo/internal/infrastructure/database"
	"hiyoko-echo/internal/interactor"
)

const (
	EnvRoot         = "cmd/cli"
	DBQueryPing     = "ping"
	DBQueryMigrate  = "migrate"
	DBQuerySeed     = "seed"
	DBQueryTruncate = "truncate"
	DBQueryDrop     = "drop"

	ErrDefaultMsg      = "failed to query %s; error: %v"
	QuerySuccessfulMsg = "success query %s"
)

var (
	serverEnv    util.ServerEnv
	databaseConf database.Conf
	query        *string
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
	query = flag.String("query", "ping", "exec query")
	flag.Parse()

	// load env
	serverEnv = util.ServerEnv(*server)
	if ok := serverEnv.Regexp(); !ok {
		panic("invalid server environment")
	}
	util.LoadEnv(serverEnv, EnvRoot)
	databaseConf = configs.NewMySqlConf()

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

	ctx := context.Background()
	i := interactor.NewInteractor(entClient)
	r := i.NewTableRepository()

	switch *query {
	case DBQueryPing:
		err := r.Ping(ctx)
		if err != nil {
			log.Panicf(ErrDefaultMsg, DBQueryPing, err)
		}
		fmt.Printf(QuerySuccessfulMsg, DBQueryPing)
	case DBQueryMigrate:
		err := r.Migrate(ctx)
		if err != nil {
			log.Panicf(ErrDefaultMsg, DBQueryMigrate, err)
		}
		fmt.Printf(QuerySuccessfulMsg, DBQueryMigrate)
	case DBQuerySeed:
		err := r.Seed(ctx)
		if err != nil {
			log.Panicf(ErrDefaultMsg, DBQuerySeed, err)
		}
		fmt.Printf(QuerySuccessfulMsg, DBQuerySeed)
	case DBQueryTruncate:
		err := r.TruncateAll(ctx)
		if err != nil {
			log.Panicf(ErrDefaultMsg, DBQueryTruncate, err)
		}
		fmt.Printf(QuerySuccessfulMsg, DBQueryTruncate)
	case DBQueryDrop:
		err := r.DropAll(ctx)
		if err != nil {
			log.Panicf(ErrDefaultMsg, DBQueryDrop, err)
		}
		fmt.Printf(QuerySuccessfulMsg, DBQueryDrop)
	}
}
