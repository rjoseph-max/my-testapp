package main

import (
	"bytes"
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/maxexllc/lib-go/awsutils/apigw"
	"github.com/maxexllc/lib-go/logging"
)

// This is an example handler struct which contains S3 client
type (
	s3Client interface {
		PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	}

	handler struct {
		s3 s3Client
	}
)

// Handle is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func (h handler) Handle(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// initialize logger
	logger := logging.NewLambdaLoggerEntry(ctx, conf.LogLevel.String())
	// Set logger context for downstream utilities like 'apigw.JSON(...)' which logs response data
	ctx = logging.NewContext(ctx, logger)

	logger.WithField("body", event.Body).Info("received request")

	// unmarshal request from events.APIGatewayProxyRequest body
	var req ExampleRequest
	if err := apigw.Bind(ctx, event, &req); err != nil {
		return apigw.JSON(ctx, http.StatusBadRequest, map[string]string{"message": ErrMessageCouldNotParseRequest})
	}

	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String("test-bucket-name"),
		Key:    &req.LoanID,
		Body:   bytes.NewReader([]byte("test file data")),
	}
	if _, err := h.s3.PutObject(ctx, putObjectInput); err != nil {
		logger.WithError(err).Error("error while uploading object to S3")
		return apigw.Error(ctx, err)
	}

	return apigw.JSON(ctx, http.StatusOK, map[string]interface{}{"message": "ok"})
}
