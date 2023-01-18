package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger = initLogger()

//initLogger 初始化zap日志
func initLogger() *zap.SugaredLogger {
	zap.NewDevelopmentConfig()
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.FunctionKey = "F"
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.AddSync(os.Stderr), zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}
