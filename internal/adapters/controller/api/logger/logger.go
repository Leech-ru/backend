package logger

import (
	"Leech-ru/internal/adapters/config"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func SetupLogger(s *ghttp.Server, cfg config.LoggerConfig) {
	logger := glog.New()
	logger.SetFlags(glog.F_TIME_DATE)
	logger.Level(glog.LEVEL_ALL)

	s.SetLogger(logger)

	s.SetLogStdout(true)
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)
}
