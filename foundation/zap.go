package foundation

import (
	"fmt"
	"github.com/golang-work/adventure/support"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
	if ok, _ := support.PathExists(support.Config["app"].GetString("zap.director")); !ok {
		fmt.Printf("create %v directory\n", support.Config["app"].GetString("zap.director"))
		_ = os.Mkdir(support.Config["app"].GetString("zap.director"), os.ModePerm)
	}

	switch support.Config["app"].GetString("zap.level") {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(encoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(encoderCore())
	}
	if support.Config["app"].GetBool("zap.showLine") {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func encoderCore() (core zapcore.Core) {
	name := path.Join(support.Config["app"].GetString("zap.director"),
		fmt.Sprintf("%s.log", support.Now().Format("2006-01-02")))
	writer, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	zapcore.AddSync(writer)

	return zapcore.NewCore(encoder(), writer, level)
}

func encoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format(support.Config["app"].GetString("zap.prefix") + "2006-01-02 15:04:05"))
	}
	if support.Config["app"].GetString("zap.format") == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}
