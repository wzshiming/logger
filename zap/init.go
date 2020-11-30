package zap

import (
	"os"

	"github.com/wzshiming/logger"
)

var Log = New(WriteTo(os.Stderr), UseDevMode(true))

func init() {
	logger.SetLogger(Log)
}
