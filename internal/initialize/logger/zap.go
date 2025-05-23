package logger

import (
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"example/internal/global"
)

func Init(debug bool, options ...zap.Option) {
	var coreArr []zapcore.Core

	// 获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()            // NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder    // 指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 按级别显示不同颜色
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder       // 显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		if debug {
			return lev < zap.ErrorLevel && lev >= zap.DebugLevel
		} else {
			return lev < zap.ErrorLevel && lev > zap.DebugLevel
		}
	})

	path := strings.TrimRight(global.Config.Log.Path, "/")

	// info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path + "/info.log", // 日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    100,                // 文件大小限制,单位MB
		MaxBackups: 100,                // 最大保留日志文件数量
		MaxAge:     7,                  // 日志文件保留天数
		Compress:   false,              // 是否压缩处理
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer), lowPriority) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path + "/error.log", // 日志文件存放目录
		MaxSize:    100,                 // 文件大小限制,单位MB
		MaxBackups: 100,                 // 最大保留日志文件数量
		MaxAge:     7,                   // 日志文件保留天数
		Compress:   false,               // 是否压缩处理
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	opts := []zap.Option{zap.AddCaller()}
	if debug {
		opts = append(opts, zap.AddStacktrace(lowPriority))
	}
	opts = append(opts, options...)
	global.Logger = zap.New(zapcore.NewTee(coreArr...), opts...) // zap.AddCaller()为显示文件名和行号
}
