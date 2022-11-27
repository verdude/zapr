package zapr

import (
  "go.uber.org/zap"
)

const (
  ZSugar = iota
  ZLogger
)

type settings struct { }

type loggingT struct {
  settings
  vLevel uint8
  mode uint
}

var logging loggingT

func Sugar() {
  logging.mode = ZSugar
}

func Logger() {
  logging.mode = ZLogger
}
