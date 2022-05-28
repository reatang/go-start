package xhttp

import (
	"github.com/go-resty/resty/v2"
	"github.com/reatang/gostart/pkg/logger"
)

func requestLogCallback(log *resty.RequestLog) error {
	logger.Debug(log.Body)
	return nil
}

func responseLogCallback(log *resty.ResponseLog) error {
	logger.Debug(log.Body)
	return nil
}
