// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package space

import (
	"context"
	"encoding/json"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
	daprclt "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/client"
	ierr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/errors"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
)

type Space interface {
	AllocateSpace(ctx context.Context, req AllocateSpaceReq) (AllocateSpaceRes, error)
	DeallocateSpace(ctx context.Context, req DeallocateSpaceReq) error
}

type space struct {
	logger logger.Logger
	client daprclt.Client
}

func NewSpace(logger logger.Logger, client daprclt.Client) Space {
	return space{
		logger: logger,
		client: client,
	}
}

func (o space) AllocateSpace(ctx context.Context, req AllocateSpaceReq) (AllocateSpaceRes, error) {
	o.logger.Infof("Req data to space AllocateSpace method: %v", req)

	data, err := o.client.InvokeMethodWithContent(ctx, constant.Space, constant.AllocateSpace, req)
	if err != nil {
		return AllocateSpaceRes{}, ierr.WrapErrorf(err, ierr.Unknown, "client.InvokeMethodWithContent")
	}

	res := AllocateSpaceRes{}
	if err := json.Unmarshal(data, &res); err != nil {
		return AllocateSpaceRes{}, ierr.WrapErrorf(err, ierr.Unknown, "json.Unmarshal")
	}

	o.logger.Infof("Res data from space AllocateSpace method: %v", res)

	return res, nil
}

func (o space) DeallocateSpace(ctx context.Context, req DeallocateSpaceReq) error {
	o.logger.Infof("Req data to space DeallocateSpace method: %v", req)

	_, err := o.client.InvokeMethodWithContent(ctx, constant.Space, constant.DeallocateSpace, req)
	if err != nil {
		return ierr.WrapErrorf(err, ierr.Unknown, "client.InvokeMethodWithContent")
	}

	return nil
}
