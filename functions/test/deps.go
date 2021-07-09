package main

import (
	"context"
	"io/ioutil"

	"github.com/maxexllc/lib-go/logging"
	"github.com/sirupsen/logrus"
)

func MockCtx() context.Context {
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	return logging.NewContext(context.Background(), &logging.Entry{Entry: logrus.NewEntry(logger)})
}
