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

	"github.com/labstack/echo/v4"
	"hiyoko-echo/cmd/util"
	"hiyoko-echo/conf"
	"hiyoko-echo/infrastructure/database"
	logger "hiyoko-echo/infrastructure/logger/local"
	"hiyoko-echo/interactor"
	"hiyoko-echo/presenter/http/middleware"
	"hiyoko-echo/presenter/http/router"
	"hiyoko-echo/shared"
)

const (
	envRoot = "cmd/public"
)

var (
	loging       logger.Logger
	serverEnv    util.ServerEnv
	databaseConf database.Conf
	ctx          context.Context
)

func init() {
	ctx = context.Background()
	logFilepath, err := shared.GetLogFilePath(conf.LogPath)
	if err != nil {
		log.Fatalf("failed to get executable path; error: %v", err)
	}
	loging = logger.NewLogger(logFilepath)

	// seed
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		loging.Fatalf(ctx, "failed to create seed; error: %v", err)
	}
	rand.Seed(seed.Int64()) // go1.20から自動設定される

	// flag
	server := flag.String("server", "local", "server environment")
	flag.Parse()

	// load env
	serverEnv = util.ServerEnv(*server)
	if ok := serverEnv.Regexp(); !ok {
		loging.Fatalf(ctx, "invalid server environment")
	}
	util.LoadEnv(serverEnv, envRoot)
	databaseConf = conf.NewMySqlConf()

	// timezone
	util.LoadTimezone()
}

func main() {
	e := echo.New()
	entClient, err := database.NewMySqlConnect(serverEnv, databaseConf)
	if err != nil {
		loging.Fatalf(ctx, "failed to create dbclient; error: %v", err)
	}
	defer func(entClient *database.EntClient) {
		err := entClient.Close()
		if err != nil {
			loging.Fatalf(ctx, "failed to close dbclient; error: %v", err)
		}
	}(entClient)

	i := interactor.NewInteractor(entClient)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)
	if err := e.Start(fmt.Sprintf(":%d", shared.Env("SERVER_PORT").GetInt(8000))); err != nil {
		loging.Fatalf(ctx, "failed to start server; error: %v", err)
	}

	// todo 正常にサーバーが立ち上がったlogを作成
	//loging.Fatalf("failed to start server; error: %v", err)
}
