package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/natefinch/lumberjack.v2"
	"hiyoko-echo/conf"
	"hiyoko-echo/shared"
	"io"
	"strings"
	"time"
)

func NewMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(shared.Env("CLIENT_WEB_URL").GetString("*"), ","),
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authentication",
		},
		AllowCredentials: false,
		MaxAge:           24 * int(time.Hour),
	}))

	e.Use(middleware.RequestID())
	logPath, _ := shared.GetLogFilePath(conf.LogPath)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: conf.AccessLogFormat,
		Output: io.MultiWriter(&lumberjack.Logger{
			Filename:   logPath + "/access.log",
			MaxSize:    conf.LogSize,
			MaxBackups: conf.LogBucket,
			MaxAge:     conf.LogAge,
			Compress:   conf.LogCompress,
		}),
	}))

	// todo 独自ロガーを組み込む
	// todo 別ファイルへ移動する
	//log := logrus.New()
	//e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	//	LogURI:    true,
	//	LogStatus: true,
	//	LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
	//		log.WithFields(logrus.Fields{
	//			"URI":   values.URI,
	//			"status": values.Status,
	//		}).Info("request")
	//
	//		return nil
	//	},
	//}))
}
