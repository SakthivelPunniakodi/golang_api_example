package common

import (
	"context"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-space/internal/dto"
)

type SpaceService interface {
	AllocateSpace(ctx context.Context, req dto.AllocateSpaceReq) (dto.AllocateSpaceRes, error)
	DeallocateSpace(ctx context.Context, req dto.DeallocateSpaceReq) error
	PublishSpotChangedEvent(ctx context.Context) error
}
