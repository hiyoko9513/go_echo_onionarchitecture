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

	"hiyoko-echo/configs"
	"hiyoko-echo/internal/infrastructure/database"
	"hiyoko-echo/internal/interactor"
	logger "hiyoko-echo/internal/pkg/logger/local"
	"hiyoko-echo/internal/presentation/http/middleware"
	"hiyoko-echo/internal/presentation/http/router"
	"hiyoko-echo/util"

	"github.com/labstack/echo/v4"
)

const (
	envRoot = "cmd/public"
)

var (
	logging      logger.Logger
	serverEnv    util.ServerEnv
	databaseConf database.Conf
	ctx          context.Context
)

func init() {
	ctx = context.Background()
	logFilepath, err := util.GetLogFilePath(configs.LogPath)
	if err != nil {
		log.Fatalf("failed to get executable path; error: %v", err)
	}
	logging = logger.NewLogger(logFilepath)

	// seed
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		logging.Fatalf(ctx, "failed to create seed; error: %v", err)
	}
	rand.NewSource(seed.Int64())

	// flag
	server := flag.String("server", "local", "server environment")
	flag.Parse()

	// load env
	serverEnv = util.ServerEnv(*server)
	if ok := serverEnv.Regexp(); !ok {
		logging.Fatalf(ctx, "invalid server environment")
	}
	util.LoadEnv(serverEnv, envRoot)
	databaseConf = configs.NewMySqlConf()

	// timezone
	util.LoadTimezone()
}

func main() {
	e := echo.New()
	e.HideBanner = true
	entClient, err := database.NewMySqlConnect(serverEnv, databaseConf)
	if err != nil {
		logging.Fatalf(ctx, "failed to create dbclient; error: %v", err)
	}
	defer func(entClient *database.EntClient) {
		err := entClient.Close()
		if err != nil {
			logging.Fatalf(ctx, "failed to close dbclient; error: %v", err)
		}
	}(entClient)

	i := interactor.NewInteractor(entClient)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)
	if err := e.Start(fmt.Sprintf(":%d", util.Env("SERVER_PORT").GetInt(8000))); err != nil {
		logging.Fatalf(ctx, "failed to start server; error: %v", err)
	}

	// todo 正常にサーバーが立ち上がったlogを作成
	//logging.Fatalf("failed to start server; error: %v", err)
}
