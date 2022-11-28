package zapr

import (
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

const (
  ZSugar = iota
  ZLogger
)

type settings struct { }

type loggingT struct {
  settings
  vLevel uint8
  logger *zap.Logger
}

type bervose struct {
  enabled bool
}

func (b bervose) Enabled() bool {
  return b.enabled
}

func (b bervose) I(m string, args ...zapcore.Field) {
  if b.enabled {
    logging.logger.Info(m, args...)
  }
}

func (b bervose) W(m string, args ...zapcore.Field) {
  if b.enabled {
    logging.logger.Warn(m, args...)
  }
}

func (b bervose) E(m string, args ...zapcore.Field) {
  if b.enabled {
    logging.logger.Error(m, args...)
  }
}

func (b bervose) D(m string, args ...zapcore.Field) {
  if b.enabled {
    logging.logger.Debug(m, args...)
  }
}

var logging loggingT

// API
func I(m string, args ...zapcore.Field) {
  logging.logger.Info(m, args...)
}

func E(m string, args ...zapcore.Field) {
  logging.logger.Error(m, args...)
}

func D(m string, args ...zapcore.Field) {
  logging.logger.Debug(m, args...)
}

func W(m string, args ...zapcore.Field) {
  logging.logger.Warn(m, args...)
}

func Sync() {
  logging.logger.Sync()
}

func V(level uint8) *bervose {
  return &bervose{level <= logging.vLevel}
}
