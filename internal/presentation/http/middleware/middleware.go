package middleware

import (
	"io"
	"strings"
	"time"

	"hiyoko-echo/configs"
	"hiyoko-echo/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(util.Env("CLIENT_WEB_URL").GetString("*"), ","),
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

	logPath, _ := util.GetLogFilePath(configs.LogPath)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: configs.AccessLogFormat,
		Output: io.MultiWriter(&lumberjack.Logger{
			Filename:   logPath + "/access.log",
			MaxSize:    configs.LogSize,
			MaxBackups: configs.LogBucket,
			MaxAge:     configs.LogAge,
			Compress:   configs.LogCompress,
		}),
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqID := c.Response().Header().Get(echo.HeaderXRequestID)
			c.Set("RequestID", reqID)
			return next(c)
		}
	})
}
