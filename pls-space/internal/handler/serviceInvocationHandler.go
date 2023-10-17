// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"context"
	"encoding/json"
	"fmt"

	ierr "sampleapi/pls-shared/errors"
	"sampleapi/pls-space/internal/dto"
)

type serviceInvocationHandler interface {
	AllocateSpace(ctx context.Context, data []byte) (any, error)
	DeallocateSpace(ctx context.Context, data []byte) (any, error)
}

func (o spaceHandler) AllocateSpace(ctx context.Context, data []byte) (any, error) {
	req := dto.AllocateSpaceReq{}
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, ierr.WrapErrorf(err, ierr.Unknown, "json.Unmarshal")
	}

	o.logger.Infof("Req data to space AllocateSpace method: %v", req)

	res, err := o.spaceSvc.AllocateSpace(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("spaceSvc.AllocateSpace: %w", err)
	}

	return res, nil
}

func (o spaceHandler) DeallocateSpace(ctx context.Context, data []byte) (any, error) {
	req := dto.DeallocateSpaceReq{}
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, ierr.WrapErrorf(err, ierr.Unknown, "json.Unmarshal")
	}

	o.logger.Infof("Req data to space DeallocateSpace method: %v", req)

	if err := o.spaceSvc.DeallocateSpace(ctx, req); err != nil {
		return nil, fmt.Errorf("spaceSvc.DeallocateSpace: %w", err)
	}

	return nil, nil
}
