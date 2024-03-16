package main

import (
	"context"
	crand "crypto/rand"
	"flag"
	"math"
	"math/big"
	"math/rand"

	"hiyoko-echo/configs"
	"hiyoko-echo/internal/infrastructure/database"
	"hiyoko-echo/internal/interactor"
	"hiyoko-echo/pkg/logging/file"
	"hiyoko-echo/util"
)

const (
	EnvRoot         = "cmd/cli"
	DBQueryPing     = "ping"
	DBQueryMigrate  = "migrate"
	DBQuerySeed     = "seed"
	DBQueryTruncate = "truncate"
	DBQueryDrop     = "drop"

	ErrDefaultMsg      = "failed to query"
	QuerySuccessfulMsg = "success query"
)

const logDir = "./log/cli/db"

var (
	serverEnv    util.ServerEnv
	databaseConf database.Conf
	query        *string
)

func init() {
	// flag
	server := flag.String("server", "local", "server environment")
	query = flag.String("query", "ping", "exec query")
	flag.Parse()

	logger.SetLogDir(logDir)
	logger.Initialize()
	logger.With("query", query)

	// seed
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		logger.Fatal("failed to create seed", "error", err)
	}
	rand.NewSource(seed.Int64())

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
		logger.Fatal("failed to create dbclient", "error", err)
	}
	defer func(entClient *database.EntClient) {
		err := entClient.Close()
		if err != nil {
			logger.Fatal("failed to close dbclient", "error", err)
		}
	}(entClient)

	ctx := context.Background()
	i := interactor.NewInteractor(entClient)
	r := i.NewTableRepository()

	switch *query {
	case DBQueryPing:
		err := r.Ping(ctx)
		if err != nil {
			logger.Fatal(ErrDefaultMsg, "error", err)
		}
	case DBQueryMigrate:
		err := r.Migrate(ctx)
		if err != nil {
			logger.Fatal(ErrDefaultMsg, "error", err)
		}
	case DBQuerySeed:
		err := r.Seed(ctx)
		if err != nil {
			logger.Fatal(ErrDefaultMsg, "error", err)
		}
	case DBQueryTruncate:
		err := r.TruncateAll(ctx)
		if err != nil {
			logger.Fatal(ErrDefaultMsg, "error", err)
		}
	case DBQueryDrop:
		err := r.DropAll(ctx)
		if err != nil {
			logger.Fatal(ErrDefaultMsg, "error", err)
		}
		logger.Info(QuerySuccessfulMsg)
	}
}
