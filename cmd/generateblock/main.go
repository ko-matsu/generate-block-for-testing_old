package main

import (
	"context"
	"fmt"

	arg "github.com/alexflint/go-arg"
	env "github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func logInit(argObj *argument) *zap.Logger {
	logConf := zap.NewProductionConfig()
	logOpts := []zap.Option{zap.AddStacktrace(zapcore.WarnLevel)}
	if argObj.Logging {
		logConf.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		logConf.Development = true
	}
	logger, err := logConf.Build(logOpts...)
	if err != nil {
		panic(err)
	}
	return logger
}

func main() {
	ctx := context.Background()
	argObj := &argument{}
	arg.MustParse(argObj)

	logger := logInit(argObj)
	logger.Debug("start")

	envObj := &environment{}
	if err := env.Parse(envObj); err != nil {
		if argObj.Logging {
			logger.Error("Error while parsing environment", zap.Error(err))
		}
		panic(err)
	}

	argObj.setValueFromEnvironment(envObj)
	if err := argObj.Validate(); err != nil {
		if argObj.Logging {
			logger.Error("Error while validate argument", zap.Error(err))
		} else {
			fmt.Printf("ERROR: %s\n", err.Error())
		}
		return
	}
	config := argObj.ToConfigurationModel()

	// dedpendency
	handle := NewHandler(argObj)

	// execute
	if err := handle.GenerateBlock(ctx, config); err != nil {
		if argObj.Logging {
			logger.Error("GenerateBlock fail", zap.Error(err))
		} else {
			fmt.Printf("ERROR: %s\n", err.Error())
		}
	}

	logger.Debug("end")
}
