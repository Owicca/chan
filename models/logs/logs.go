package logs

import (
	"time"

	"upspin.io/errors"
	"go.uber.org/zap"
)

// Use upspin error as a formatter
// and send it to zap logger.
func LogInfo(params ...interface{}) error {
	err := errors.E(params...)
	zap.L().Info(err.Error(), zap.Int64("timestamp", time.Now().Unix()))

	return err
}

// Use upspin error as a formatter
// and send it to zap logger.
func LogWarn(params ...interface{}) error {
	err := errors.E(params...)
	zap.L().Warn(err.Error(), zap.Int64("timestamp", time.Now().Unix()))

	return err
}

// Use upspin error as a formatter
// and send it to zap logger.
func LogErr(params ...interface{}) error {
	err := errors.E(params...)
	zap.L().Error(err.Error(), zap.Int64("timestamp", time.Now().Unix()))

	return err
}