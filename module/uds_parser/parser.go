package uds_parser

import (
	"context"
	uds "explain_uds/common"
)

func ParseUDS(ctx context.Context, raw []byte) (*UDSMapping, error) {
	if len(raw) < 1 {
		return nil, uds.ErrEmptyMessage
	}
	serviceId := raw[0]
	serviceDTO, err := GetServiceByID(ctx, string(serviceId))
	if err != nil {
		return nil, err
	}
	if serviceDTO.HasSubFunction {

	}
	res := &UDSMapping{
		SID:         serviceDTO.SID,
		Name:        serviceDTO.Name,
		Description: serviceDTO.Description,
	}
	return res, nil
}
