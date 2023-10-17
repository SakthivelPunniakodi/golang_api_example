// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"github.com/SakthivelPunniakodi/golang_api_example/pls-manager/internal/service"
	spaceEvents "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/events/space"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"

	"github.com/gorilla/websocket"
)

type Option func(*manager)

func WithLogger(logger logger.Logger) Option {
	return func(o *manager) {
		o.logger = logger
	}
}

func WithSvc(managerSvc service.Manager) Option {
	return func(o *manager) {
		o.managerSvc = managerSvc
	}
}

func WithUnoccupiedSpotsChannel(unoccupiedSpotsCh <-chan []spaceEvents.Floor) Option {
	return func(o *manager) {
		o.unoccupiedSpotsCh = unoccupiedSpotsCh
	}
}

type Manager interface {
	externalHandler
	topicEventHandler
}

type manager struct {
	logger            logger.Logger
	managerSvc        service.Manager
	unoccupiedSpotsCh <-chan []spaceEvents.Floor
	websocketUpgrader websocket.Upgrader
}

// NewManager handler instance.
func NewManager(opts ...Option) Manager {
	o := &manager{
		websocketUpgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
