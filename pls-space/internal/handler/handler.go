// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-space/internal/common"
)

type Option func(*spaceHandler)

func WithLogger(logger logger.Logger) Option {
	return func(o *spaceHandler) {
		o.logger = logger
	}
}

func WithSvc(spaceSvc common.SpaceService) Option {
	return func(o *spaceHandler) {
		o.spaceSvc = spaceSvc
	}
}

type SpaceHandler interface {
	serviceInvocationHandler
}

type spaceHandler struct {
	logger   logger.Logger
	spaceSvc common.SpaceService
}

// NewSpaceHandler handler instance.
func NewSpaceHandler(opts ...Option) SpaceHandler {
	o := &spaceHandler{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
