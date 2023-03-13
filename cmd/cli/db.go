package main

import (
	"context"
	crand "crypto/rand"
	"flag"
	"fmt"
	"hiyoko-echo/interactor"
	"math"
	"math/big"
	"math/rand"

	"hiyoko-echo/cmd/util"
	"hiyoko-echo/conf"
	"hiyoko-echo/infrastructure/database"
)

var (
	serverEnv    util.ServerEnv
	databaseConf database.Conf
	migrateFlg   *bool
	seedFlg      *bool
	truncateFlg  *bool
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
	migrateFlg = flag.Bool("migrate", false, "run db migrations")
	seedFlg = flag.Bool("seed", false, "run db seeding")
	truncateFlg = flag.Bool("truncate", false, "run db truncate")
	flag.Parse()

	// load env
	serverEnv = util.ServerEnv(*server)
	if ok := serverEnv.Regexp(); !ok {
		panic("invalid server environment")
	}
	util.LoadEnv(serverEnv, "cmd/cli")
	databaseConf = conf.NewMySqlConf()

	// timezone
	util.LoadTimezone()
}

// entのauto migrate機能を利用
func main() {
	entClient, err := database.NewMySqlConnect(serverEnv, databaseConf)
	if err != nil {
		panic(err)
	}
	defer func(entClient *database.EntClient) {
		err := entClient.Close()
		if err != nil {
			panic(err)
		}
	}(entClient)

	ctx := context.Background()
	i := interactor.NewInteractor(entClient)
	r := i.NewTableRepository()

	// todo 以下あまりよくないので、修正する
	if *migrateFlg {
		err := r.Migrate(ctx)
		if err != nil {
			panic(err)
		}
	}

	// todo seederの作成
	if *seedFlg {
		fmt.Println("seed test")
	}

	if *truncateFlg {
		err := r.TruncateAll(ctx)
		if err != nil {
			panic(err)
		}
	}
}
