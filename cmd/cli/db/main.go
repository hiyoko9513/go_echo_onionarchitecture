package main

import (
	"context"
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
	"hiyoko-echo/interactor"
)

const (
	EnvRoot         = "cmd/cli"
	DBQueryPing     = "ping"
	DBQueryMigrate  = "migrate"
	DBQuerySeed     = "seed"
	DBQueryTruncate = "truncate"
	DBQueryDrop     = "drop"
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
	rand.Seed(seed.Int64()) // go1.20から自動設定される

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

	ctx := context.Background()
	i := interactor.NewInteractor(entClient)
	r := i.NewTableRepository()

	// todo あんまりいい作りではないような？
	switch *query {
	case DBQueryPing:
		err := r.Ping(ctx)
		if err != nil {
			log.Panicf("failed to query %s; error: %v", DBQueryPing, err)
		}
		fmt.Printf("success query %s", DBQueryPing)
	case DBQueryMigrate:
		err := r.Migrate(ctx)
		if err != nil {
			log.Panicf("failed to query %s; error: %v", DBQueryMigrate, err)
		}
		fmt.Printf("success query %s", DBQueryMigrate)
	case DBQuerySeed:
		err := r.Seed(ctx)
		if err != nil {
			log.Panicf("failed to query %s; error: %v", DBQuerySeed, err)
		}
		fmt.Printf("success query %s", DBQuerySeed)
	case DBQueryTruncate:
		err := r.TruncateAll(ctx)
		if err != nil {
			log.Panicf("failed to query %s; error: %v", DBQueryTruncate, err)
		}
		fmt.Printf("success query %s", DBQueryTruncate)
	case DBQueryDrop:
		err := r.DropAll(ctx)
		if err != nil {
			log.Panicf("failed to query %s; error: %v", DBQueryDrop, err)
		}
		fmt.Printf("success query %s", DBQueryDrop)
	}
}
