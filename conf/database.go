package conf

import (
	"hiyoko-echo/infrastructure/database"
	"hiyoko-echo/shared"
)

func NewMySqlConf() (conf database.Conf) {
	conf.Host = shared.Env("DB_HOST").GetString("localhost")
	conf.User = shared.Env("DB_USER").GetString("hiyoko")
	conf.Password = shared.Env("DB_PASSWORD").GetString("hiyoko")
	conf.Name = shared.Env("DB_NAME").GetString("hiyoko")
	conf.Port = shared.Env("DB_PORT").GetInt(3306)
	return
}
