package local

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
	"hiyoko-echo/conf"
)

// todo ログのフォーマットを指定する
// todo ログレベルを設定する
// todo ログの出力方法を設定する 基本同期的に処理したほうが良さそう

const (
	DebugFilePath = "/debug.log"
	InfoFilePath  = "/info.log"
	WarnFilePath  = "/warn.log"
	ErrorFilePath = "/error.log"
)

func init() {
	// make file and folder
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("failed to get executable path; error: %v", err)
	}

	for _, path := range conf.LogPaths {
		logPath := filepath.Join(filepath.Dir(exePath), path)
		if err = makeDirAll(logPath); err != nil {
			log.Fatalf("failed to make directory; error: %v", err)
		}
	}
}

type logger struct {
	aLog *slog.Logger
	dLog *slog.Logger
	iLog *slog.Logger
	wLog *slog.Logger
	eLog *slog.Logger
}

func NewLogger(logFilepath string) Logger {
	debugLogger := slog.New(slog.HandlerOptions{
		Level: slog.LevelDebug,
	}.NewJSONHandler(&lumberjack.Logger{
		Filename:   logFilepath + DebugFilePath,
		MaxSize:    conf.LogSize,
		MaxBackups: conf.LogBucket,
		MaxAge:     conf.LogAge,
		Compress:   conf.LogCompress,
	}))

	infoLogger := slog.New(slog.HandlerOptions{
		Level: slog.LevelInfo,
	}.NewJSONHandler(&lumberjack.Logger{
		Filename:   logFilepath + InfoFilePath,
		MaxSize:    conf.LogSize,
		MaxBackups: conf.LogBucket,
		MaxAge:     conf.LogAge,
		Compress:   conf.LogCompress,
	}))

	warnLogger := slog.New(slog.HandlerOptions{
		Level: slog.LevelWarn,
	}.NewJSONHandler(&lumberjack.Logger{
		Filename:   logFilepath + WarnFilePath,
		MaxSize:    conf.LogSize,
		MaxBackups: conf.LogBucket,
		MaxAge:     conf.LogAge,
		Compress:   conf.LogCompress,
	}))

	errLogger := slog.New(slog.HandlerOptions{
		Level: slog.LevelError,
	}.NewJSONHandler(&lumberjack.Logger{
		Filename:   logFilepath + ErrorFilePath,
		MaxSize:    conf.LogSize,
		MaxBackups: conf.LogBucket,
		MaxAge:     conf.LogAge,
		Compress:   conf.LogCompress,
	}))

	return &logger{
		dLog: debugLogger,
		iLog: infoLogger,
		wLog: warnLogger,
		eLog: errLogger,
	}
}

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
}

func (l *logger) Debug(args ...interface{}) {
	//l.log.Printf("[INFO] "+format, args...)
}
func (l *logger) Info(args ...interface{}) {
	//l.log.Printf("[INFO] "+format, args...)
}
func (l *logger) Warning(args ...interface{}) {
	//l.log.Printf("[INFO] "+format, args...)
}
func (l *logger) Error(args ...interface{}) {
	//l.log.Printf("[INFO] "+format, args...)
}
func (l *logger) Fatalf(ctx context.Context, msg string, args ...interface{}) {
	// todo フォーマットの設定
	l.eLog.ErrorCtx(ctx, msg, args)
	// todo リカバリ処理が必要？
	panic("【PANIC】 msg:" + msg + " args:" + fmt.Sprint(args...))
}

func makeDirAll(path string) error {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		defaultUmask := syscall.Umask(0)
		err := os.MkdirAll(path, 0755)
		syscall.Umask(defaultUmask)
		return err
	}
	return nil
}
