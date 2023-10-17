// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package client

import (
	"context"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
	ierr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/errors"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
)

type Option func(*client)

func WithLogger(logger logger.Logger) Option {
	return func(o *client) {
		o.logger = logger
	}
}

func WithPubSubAndTopic(pubSubName, topicName string) Option {
	return func(o *client) {
		o.pubSubName = pubSubName
		o.topicName = topicName
	}
}

type Client interface {
	InvokeMethod(ctx context.Context, appID, methodName string) ([]byte, error)
	InvokeMethodWithContent(ctx context.Context, appID, methodName string, content any) ([]byte, error)
	PublishEvent(ctx context.Context, event any) error
	Close()
}

type client struct {
	client     dapr.Client
	logger     logger.Logger
	pubSubName string
	topicName  string
}

func NewClient(opts ...Option) (Client, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, ierr.WrapErrorf(err, ierr.Unknown, "daprClient.NewClient")
	}

	return newClient(client, opts), nil
}

func newClient(c dapr.Client, opts []Option) Client {
	o := &client{client: c}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (o client) InvokeMethod(ctx context.Context, appID, methodName string) ([]byte, error) {
	res, err := o.client.InvokeMethod(ctx, appID, methodName, http.MethodPost)
	if err != nil {
		return nil, ierr.WrapErrorf(err, ierr.Unknown, "client.InvokeMethod")
	}

	return res, nil
}

func (o client) InvokeMethodWithContent(ctx context.Context, appID, methodName string, data any) ([]byte, error) {
	res, err := o.client.InvokeMethodWithCustomContent(ctx, appID, methodName, http.MethodPost, constant.ContentTypeJSON, data)
	if err != nil {
		return nil, ierr.WrapErrorf(err, ierr.Unknown, "client.InvokeMethodWithCustomContent")
	}

	return res, nil
}

func (o client) PublishEvent(ctx context.Context, event any) error {
	if err := o.client.PublishEvent(ctx, o.pubSubName, o.topicName, event); err != nil {
		return ierr.WrapErrorf(err, ierr.Unknown, "client.PublishEvent")
	}

	o.logger.Infof("Published event: %v", event)

	return nil
}

func (o client) Close() {
	o.client.Close()
}
