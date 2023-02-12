package netconnect

import "go.uber.org/zap"

type Logger struct {
	Name string
	Log  *zap.SugaredLogger
}

func NewLogger(name string) (*Logger, error) {

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &Logger{
		Name: name,
		Log:  logger.Sugar(),
	}, nil
}
