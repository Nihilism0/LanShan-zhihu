package boot

import (
	g "CSAwork/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func LoggerSetup() {
	dynamicLevel := zap.NewAtomicLevel()

	switch g.Config.Logger.LogLevel {
	case "debug":
		dynamicLevel.SetLevel(zap.DebugLevel)
	case "info":
		dynamicLevel.SetLevel(zap.InfoLevel)
	case "warn":
		dynamicLevel.SetLevel(zap.WarnLevel)
	case "error":
		dynamicLevel.SetLevel(zap.ErrorLevel)
	}

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "message",                        // 日志消息键
		LevelKey:       "level",                          // 日志等级键
		TimeKey:        "time",                           // 时间键
		NameKey:        "logger",                         // 日志记录器名
		CallerKey:      "caller",                         // 日志文件信息键
		StacktraceKey:  "stacktrace",                     // 堆栈键
		LineEnding:     zapcore.DefaultLineEnding,        // 友好日志换行符
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 友好日志等级名大小写（info INFO）
		EncodeTime:     CustomTimeEncoder,                // 友好日志时日期格式化
		EncodeDuration: zapcore.StringDurationEncoder,    // 时间序列化
		EncodeCaller:   zapcore.FullCallerEncoder,        // 日志文件信息 short（包/文件.go:行号） full (文件位置.go:行号)
	})
	cores := [...]zapcore.Core{
		zapcore.NewCore(encoder, os.Stdout, dynamicLevel),
		zapcore.NewCore(
			encoder,
			getWriteSyncer(g.Config.Logger.SavePath),
			dynamicLevel,
		),
	}
	g.Logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	defer func(Logger *zap.Logger) {
		_ = Logger.Sync()
	}(g.Logger)
	g.Logger.Info("initialize logger successfully!")
}

// CustomTimeEncoder 格式化时间
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05.000]"))
}

func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,                       // 日志保存位置
		MaxSize:    g.Config.Logger.MaxSize,    // 日志文件最大大小 (MB)
		MaxBackups: g.Config.Logger.MaxBackups, // 日志文件备份数量
		MaxAge:     g.Config.Logger.MaxAge,     // 日志保存时间
		Compress:   g.Config.Logger.IsCompress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
