package main

import (
	crand "crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"math/rand"

	"github.com/labstack/echo/v4"
	"hiyoko-echo/cmd/util"
	"hiyoko-echo/conf"
	"hiyoko-echo/infrastructure/database"
	"hiyoko-echo/interactor"
	"hiyoko-echo/presenter/http/middleware"
	"hiyoko-echo/presenter/http/router"
	"hiyoko-echo/shared"
)

var (
	serverEnv    util.ServerEnv
	databaseConf database.Conf
)

func init() {
	// seed
	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	rand.Seed(seed.Int64()) // go1.20から自動設定される

	// flag
	server := flag.String("server", "local", "server environment")
	flag.Parse()

	// load env
	serverEnv = util.ServerEnv(*server)
	if ok := serverEnv.Regexp(); !ok {
		panic("invalid server environment")
	}
	util.LoadEnv(serverEnv, "cmd/public")
	databaseConf = conf.NewMySqlConf()

	// timezone
	util.LoadTimezone()
}

func main() {
	e := echo.New()
	entClient, err := database.NewMySqlConnect(serverEnv, databaseConf)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer func(entClient *database.EntClient) {
		err := entClient.Close()
		if err != nil {
			e.Logger.Fatal(err)
		}
	}(entClient)

	i := interactor.NewInteractor(entClient)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)
	if err := e.Start(fmt.Sprintf(":%d", shared.Env("SERVER_PORT").GetInt(8000))); err != nil {
		e.Logger.Fatal(err)
	}
}
