package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"time"
)

type Log struct {
	l *zap.Logger
}

func (l *Log) init(config *LogConfig) {
	// 实例化 zap
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 实现三个判断日志等级的interface
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	// 获取 info、warn、debug日志文件的io.Writer 抽象 SetLogFile() 在下方实现
	infoWriter := l.SetLogFile(config.Path + config.InfoFile, config.Format)
	errorWriter := l.SetLogFile(config.Path + config.InfoFile, config.Format)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)

	// 开启文件及行号
	development := zap.Development()
	l.l = zap.New(core, zap.AddCaller(), development)
}

// SetLogFile 设置日志切割方式方式
func (l *Log) SetLogFile(filename, format string) io.Writer {
	hook, err := rotatelogs.New(
		filename + format + ".log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		log.Fatal(err)
	}

	return hook
}