package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
)

type Config struct {
	LogLevel logrus.Level `envconfig:"LOGLEVEL" default:"info"`
}

var (
	conf Config
)

func main() {
	// Load AWS service configuration
	awsCfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	if err := envconfig.Process("", &conf); err != nil {
		log.Fatal(err)
	}

	h := handler{
		s3: s3.NewFromConfig(awsCfg),
	}
	lambda.Start(h.Handle)
}
