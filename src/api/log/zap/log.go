 package zap

 import (
	 "go.uber.org/zap"
	 "go.uber.org/zap/zapcore"
 )

var (
	Log *zap.Logger
)

func init() {

	logConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
	}

	Logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	Log = Logger

}
func Field(key string, value interface{}) zap.Field	 {
	return zap.Any(key, value)
}

func Info(msg string, fields ...zap.Field) {
	// fields are extra fields added to the log output
	Log.Info(msg, fields...)
	Log.Sync()
}
