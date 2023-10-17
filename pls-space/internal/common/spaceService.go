package common

import (
	"context"

	"sampleapi/pls-space/internal/dto"
)

type SpaceService interface {
	AllocateSpace(ctx context.Context, req dto.AllocateSpaceReq) (dto.AllocateSpaceRes, error)
	DeallocateSpace(ctx context.Context, req dto.DeallocateSpaceReq) error
	PublishSpotChangedEvent(ctx context.Context) error
}
