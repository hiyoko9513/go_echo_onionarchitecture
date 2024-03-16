package local

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"hiyoko-echo/configs"
	"hiyoko-echo/internal/pkg/logging"

	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DebugFilePath = "/debug.log"
	InfoFilePath  = "/info.log"
	WarnFilePath  = "/warn.log"
	ErrorFilePath = "/error.log"

	RequestIDLogFormat = "ReqID:%s %s"
)

func init() {
	// make file and folder
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("failed to get executable path; error: %v", err)
	}

	for _, path := range configs.LogPaths {
		logPath := filepath.Join(filepath.Dir(exePath), path)
		if err = makeDirAll(logPath); err != nil {
			log.Fatalf("failed to make directory; error: %v", err)
		}
	}
}

type logger struct {
	iLog *slog.Logger
	wLog *slog.Logger
	eLog *slog.Logger
	dLog *slog.Logger
}

func NewLogger(logFilepath string) Logger {
	iLog := slog.New(slog.HandlerOptions{
		Level: slog.LevelInfo,
	}.NewJSONHandler(&lumberjack.Logger{
		Filename:   filepath.Join(logFilepath, InfoFilePath),
		MaxSize:    configs.LogSize,
		MaxBackups: configs.LogBucket,
		MaxAge:     configs.LogAge,
		Compress:   configs.LogCompress,
	}))

	wLog := slog.New(slog.HandlerOptions{
		Level: slog.LevelWarn,
	}.NewJSONHandler(&lumberjack.Logger{
		Filename:   filepath.Join(logFilepath, WarnFilePath),
		MaxSize:    configs.LogSize,
		MaxBackups: configs.LogBucket,
		MaxAge:     configs.LogAge,
		Compress:   configs.LogCompress,
	}))

	eLog := slog.New(slog.HandlerOptions{
		Level: slog.LevelError,
	}.NewJSONHandler(&lumberjack.Logger{
		Filename:   filepath.Join(logFilepath, ErrorFilePath),
		MaxSize:    configs.LogSize,
		MaxBackups: configs.LogBucket,
		MaxAge:     configs.LogAge,
		Compress:   configs.LogCompress,
	}))

	dLog := slog.New(slog.HandlerOptions{
		Level: slog.LevelDebug,
	}.NewJSONHandler(&lumberjack.Logger{
		Filename:   filepath.Join(logFilepath, DebugFilePath),
		MaxSize:    configs.LogSize,
		MaxBackups: configs.LogBucket,
		MaxAge:     configs.LogAge,
		Compress:   configs.LogCompress,
	}))

	return &logger{
		iLog: iLog,
		wLog: wLog,
		eLog: eLog,
		dLog: dLog,
	}
}

type Logger interface {
	Info(ctx context.Context, msg string, args ...interface{})
	Warning(ctx context.Context, msg string, args ...interface{})
	Error(ctx context.Context, msg string, args ...interface{})
	Fatalf(ctx context.Context, msg string, args ...interface{})
	Debug(ctx context.Context, msg string, args ...interface{})
}

func (l *logger) Info(ctx context.Context, msg string, args ...interface{}) {
	reqID := logging.GetRequestIDFromContext(ctx)
	l.iLog.InfoCtx(ctx, fmt.Sprintf(RequestIDLogFormat, reqID, msg), args...)
}
func (l *logger) Warning(ctx context.Context, msg string, args ...interface{}) {
	reqID := logging.GetRequestIDFromContext(ctx)
	l.wLog.WarnCtx(ctx, fmt.Sprintf(RequestIDLogFormat, reqID, msg), args...)
}
func (l *logger) Error(ctx context.Context, msg string, args ...interface{}) {
	reqID := logging.GetRequestIDFromContext(ctx)
	l.eLog.ErrorCtx(ctx, fmt.Sprintf(RequestIDLogFormat, reqID, msg), args...)
}

// Fatalf exit application
func (l *logger) Fatalf(ctx context.Context, msg string, args ...interface{}) {
	reqID := logging.GetRequestIDFromContext(ctx)
	l.eLog.ErrorCtx(ctx, fmt.Sprintf(RequestIDLogFormat, reqID, msg), args...)
	os.Exit(1)
}
func (l *logger) Debug(ctx context.Context, msg string, args ...interface{}) {
	reqID := logging.GetRequestIDFromContext(ctx)
	l.dLog.DebugCtx(ctx, fmt.Sprintf(RequestIDLogFormat, reqID, msg), args...)
}

func makeDirAll(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		defaultUmask := syscall.Umask(0)
		err := os.MkdirAll(path, 0755)
		syscall.Umask(defaultUmask)
		if err != nil {
			return fmt.Errorf("failed to make directory; error: %w", err)
		}
	}
	return nil
}
