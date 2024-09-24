package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

// InitLogger 初始化Zap日志，输出到文件并打印到控制台
func InitLogger() {
	// 配置lumberjack，用于日志滚动和文件管理
	fileLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/server.log", // 日志文件路径
		MaxSize:    100,              // 每个日志文件的最大尺寸 (MB)
		MaxBackups: 5,                // 保留最多5个旧日志文件
		MaxAge:     30,               // 日志最多保存30天
		Compress:   true,             // 启用日志压缩
	})

	// 设置日志编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"                             // 时间键
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder      // 使用ISO8601时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder    // 日志级别大写
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder    // 简短的文件路径
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig) // 控制台格式化日志
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)       // 文件记录日志为JSON格式

	// 创建控制台输出
	consoleLogger := zapcore.AddSync(os.Stdout)

	// 创建Zap Core，多路输出到控制台和文件
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleLogger, zapcore.InfoLevel), // 控制台日志
		zapcore.NewCore(fileEncoder, fileLogger, zapcore.InfoLevel),       // 文件日志
	)

	// 创建logger
	logger = zap.New(core, zap.AddCaller()) // AddCaller用于输出调用日志的文件和行号
	logger.Info("Logger initialized successfully with file rotation and console output")
}

// LogInfo 记录普通信息日志
func LogInfo(msg string) {
	logger.Info(msg)
}

// LogError 记录错误信息日志
func LogError(err error) {
	logger.Error("Error occurred", zap.Error(err))
}
