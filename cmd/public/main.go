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
	"hiyoko-echo/internal/interactor"
	"hiyoko-echo/internal/presentation/http/middleware"
	"hiyoko-echo/internal/presentation/http/router"
	"hiyoko-echo/pkg/logging/file"
	"hiyoko-echo/util"

	"github.com/labstack/echo/v4"
)

const (
	envRoot = "cmd/public"
	logDir  = "./log/public"
)

var (
	serverEnv    util.ServerEnv
	databaseConf database.Conf
)

func init() {
	logger.SetLogDir(logDir)
	logger.Initialize()

	// seed
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		logger.Fatal("failed to create seed", "error", err)
	}
	rand.NewSource(seed.Int64())

	// flag
	server := flag.String("server", "local", "server environment")
	flag.Parse()

	// load env
	serverEnv = util.ServerEnv(*server)
	if ok := serverEnv.Regexp(); !ok {
		logger.Fatal("invalid server environment")
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
		logger.Fatal("failed to create dbclient", "error", err)
	}
	defer func(entClient *database.EntClient) {
		err := entClient.Close()
		if err != nil {
			logger.Fatal("failed to close dbclient", "error", err)
		}
	}(entClient)

	i := interactor.NewInteractor(entClient)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)
	if err := e.Start(fmt.Sprintf(":%d", util.Env("SERVER_PORT").GetInt(8000))); err != nil {
		logger.Fatal("failed to start server; error", "error", err)
	}

	logger.Fatal(fmt.Sprintf("Server started on port: %d", util.Env("SERVER_PORT").GetInt(8000)))
}
