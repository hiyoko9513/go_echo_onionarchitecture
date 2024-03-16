package database

import (
	"database/sql"
	"fmt"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"hiyoko-echo/cmd/util"
	"hiyoko-echo/pkg/mypubliclib/ent"
)

type Conf struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

type EntClient struct {
	*ent.Client
}

func NewMySqlConnect(env util.ServerEnv, conf Conf) (*EntClient, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name,
	)

	// sql.DB connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		err = fmt.Errorf("failed to connect to mysql; error: %v", err)
		return &EntClient{}, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	drv := entsql.OpenDB("mysql", db)

	client := ent.NewClient(ent.Driver(drv))

	// fixme: debug mode not available
	// debug mode
	//if env != "staging" && env != "production" {
	//	client = client.Debug()
	//}

	return &EntClient{client}, nil
}
