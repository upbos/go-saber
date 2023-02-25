package log

import (
	"fmt"
	zlog "github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

var zl zlog.Logger

// Logger 系统日志配置
type Logger struct {
	Level   string       `yaml:"level"` // 日志级别
	Console bool         `yaml:"console"`
	File    *RollingFile `yaml:"file"` // 文件保存配置
}

type RollingFile struct {
	FileName   string `yaml:"file-name"`   // 文件路径
	MaxSize    int    `yaml:"max-size"`    // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int    `yaml:"max-backups"` // 日志文件最多保存多少个备份
	MaxAge     int    `yaml:"max-age"`     // 文件最多保存多少天
	Compress   bool   `yaml:"compress"`    // 是否压缩
}

func Setup(logger *Logger) {
	level, err := zlog.ParseLevel(logger.Level)
	if err != nil {
		fmt.Printf("LogLevel parsing failed: %s", err)
		level = zlog.InfoLevel
	}

	zlog.SetGlobalLevel(level)
	zlog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zlog.CallerSkipFrameCount = zlog.CallerSkipFrameCount + 1 // this line allows us to get the proper traceback for file/line number
	zlog.TimeFieldFormat = "2006-01-02 15:04:05.000"
	zlog.TimestampFunc = func() time.Time {
		return time.Now().In(time.Local)
	}

	// Initialize writer
	var writers []io.Writer
	if logger.Console {
		output := zlog.ConsoleWriter{Out: os.Stdout, TimeFormat: zlog.TimeFieldFormat}
		writers = append(writers, output)
	}
	if logger.File != nil {
		if logger.File.FileName == "" {
			fmt.Println("empty logging file name")
		}
		w := &lumberjack.Logger{
			Filename:   logger.File.FileName,
			MaxBackups: logger.File.MaxBackups,
			MaxSize:    logger.File.MaxSize,
			MaxAge:     logger.File.MaxAge,
			Compress:   logger.File.Compress,
		}
		writers = append(writers, w)
	}
	mw := io.MultiWriter(writers...)
	ctx := zlog.New(mw).With().Timestamp()

	zl = ctx.Logger()
	//l.With().Str("App", app).Str("Env", env).Int32("Shard", shard)
}

func SubLogger(module string) zlog.Logger {
	return zl.With().Str("MODULE", module).Logger()
}

func Panic(err error) {
	zl.Panic().Caller().Stack().Err(err).Send()
}

func Panicf(message string, args ...interface{}) {
	zl.Panic().Caller().Msgf(message, args...)
}

func Fatal(err error, message string) {
	zl.Fatal().Err(err).Msg(message)
}

func Fatalf(err error, message string, args ...interface{}) {
	zl.Fatal().Err(err).Msgf(message, args)
}

func Error(err error, message string) {
	zl.Error().Caller().Stack().Err(err).Msg(message)
}

func Errorf(err error, message string, args ...interface{}) {
	zl.Error().Caller().Stack().Err(err).Msgf(message, args...)
}

func Warn(message string) {
	zl.Warn().Msg(message)
}

func Warnf(message string, args ...interface{}) {
	zl.Warn().Caller().Msgf(message, args...)
}

func Info(message string) {
	zl.Info().Msg(message)
}

func Infof(message string, args ...interface{}) {
	zl.Info().Msgf(message, args...)
}

func Debug(message string) {
	zl.Debug().Msg(message)
}

func Debugf(message string, args ...interface{}) {
	zl.Debug().Msgf(message, args)
}
