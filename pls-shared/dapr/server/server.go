// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	ierr "sampleapi/pls-shared/errors"
	"sampleapi/pls-shared/logger"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

type Option func(*server)

func WithLogger(logger logger.Logger) Option {
	return func(o *server) {
		o.logger = logger
	}
}

func WithPubSub(pubSubName string) Option {
	return func(o *server) {
		o.pubSubName = pubSubName
	}
}

type ServiceInvocationHandler func(ctx context.Context, data []byte) (any, error)

type TopicEventHandler func(ctx context.Context, data []byte) (bool, error)

type Server interface {
	Start() error
	AddServiceInvocationHandler(methodName string, fn ServiceInvocationHandler) error
	AddTopicEventHandler(subscription Subscription, fn TopicEventHandler) error
}

type server struct {
	service    common.Service
	logger     logger.Logger
	pubSubName string
}

// NewServer instance.
func NewServer(port string, opts ...Option) Server {
	return newServer(daprd.NewService(fmt.Sprintf(":%s", port)), opts)
}

// NewServerWithMux instance.
func NewServerWithMux(mux *chi.Mux, port string, opts ...Option) Server {
	return newServer(daprd.NewServiceWithMux(fmt.Sprintf(":%s", port), mux), opts)
}

func newServer(service common.Service, opts []Option) Server {
	o := &server{service: service}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (o server) Start() error {
	if err := o.service.Start(); err != nil && err != http.ErrServerClosed {
		return ierr.WrapErrorf(err, ierr.Unknown, "service.Start")
	}

	return nil
}

func (o server) AddServiceInvocationHandler(methodName string, fn ServiceInvocationHandler) error {
	wrapperFn := func(fn ServiceInvocationHandler) common.ServiceInvocationHandler {
		return func(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
			res, err := fn(ctx, in.Data)
			if err != nil {
				o.logger.Errorf("service.AddServiceInvocationHandler: %v", err)
				return nil, err
			}

			data, err := json.Marshal(res)
			if err != nil {
				return nil, err
			}

			return &common.Content{
				Data:        data,
				ContentType: in.ContentType,
				DataTypeURL: in.DataTypeURL,
			}, nil
		}
	}

	if err := o.service.AddServiceInvocationHandler(methodName, wrapperFn(fn)); err != nil {
		return ierr.WrapErrorf(err, ierr.Unknown, "service.AddServiceInvocationHandler")
	}

	return nil
}

func (o server) AddTopicEventHandler(subscription Subscription, fn TopicEventHandler) error {
	wrapperFn := func(fn TopicEventHandler) common.TopicEventHandler {
		return func(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
			return fn(ctx, e.RawData)
		}
	}

	if err := o.service.AddTopicEventHandler(o.mapSubscription(subscription), wrapperFn(fn)); err != nil {
		o.logger.Errorf("service.AddTopicEventHandler: %v", err)

		return ierr.WrapErrorf(err, ierr.Unknown, "service.AddTopicEventHandler")
	}

	return nil
}

func (o server) mapSubscription(subscription Subscription) *common.Subscription {
	return &common.Subscription{
		PubsubName: o.pubSubName,
		Topic:      subscription.Topic,
		Route:      subscription.Route,
		Match:      subscription.Match,
	}
}
