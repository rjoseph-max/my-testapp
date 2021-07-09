package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
)

type mockS3Client struct {
	PutObjectOutput *s3.PutObjectOutput
	PutObjectError  error
}

func (m mockS3Client) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	return m.PutObjectOutput, m.PutObjectError
}

func TestHandler_Handle(t *testing.T) {
	testSuite := []struct {
		name         string
		request      events.APIGatewayProxyRequest
		s3Client     mockS3Client
		expectedResp events.APIGatewayProxyResponse
		expectedErr  error
	}{
		{
			name: "loan created",
			request: events.APIGatewayProxyRequest{
				Body: `{"loanID": "test","loanAmount": 1.0001}`,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			s3Client: mockS3Client{},
			expectedResp: events.APIGatewayProxyResponse{
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				StatusCode: http.StatusOK,
				Body:       `{"message":"ok"}`,
			},
		},
		{
			name: "s3 put object failure",
			request: events.APIGatewayProxyRequest{
				Body: `{"loanID": "test","loanAmount": 1.0001}`,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			s3Client: mockS3Client{
				PutObjectError: errors.New("test"),
			},
			expectedErr: errors.New("test"),
		},
		{
			name: "bad request",
			request: events.APIGatewayProxyRequest{
				Body: `{"loanID": "test","loanAmount": 1.0001}`,
				Headers: map[string]string{
					"Content-Type": "xxxxxxx",
				},
			},
			s3Client: mockS3Client{},
			expectedResp: events.APIGatewayProxyResponse{
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				StatusCode: http.StatusBadRequest,
				Body:       `{"message":"` + ErrMessageCouldNotParseRequest + `"}`,
			},
		},
	}

	for _, tc := range testSuite {
		t.Run(tc.name, func(t *testing.T) {
			h := handler{
				s3: tc.s3Client,
			}
			ctx := MockCtx()
			resp, err := h.Handle(ctx, tc.request)
			if err != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.Equal(t, tc.expectedResp, resp)
			}
		})
	}
}
