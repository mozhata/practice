package log

import "go.uber.org/zap"

var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	config := zap.NewProductionConfig()
	var err error
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}
	sugar = logger.Sugar()
}

func Sync() {
	logger.Sync()
	sugar.Sync()
}

var (
	// Infof  = sugar.Infof
	Infofw = sugar.Infow
)

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
// func (s *SugaredLogger) Infof(template string, args ...interface{}) {
// 	s.log(InfoLevel, template, args, nil)
// }
