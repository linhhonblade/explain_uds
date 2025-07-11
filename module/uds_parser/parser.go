package uds_parser

import (
	"context"
	uds "explain_uds/common"
)

func ParseUDS(ctx context.Context, message []string) (*UDSMapping, error) {
	if len(message) < 1 {
		return nil, uds.ErrEmptyMessage
	}
	serviceId := message[0]
	serviceDTO, err := GetServiceByID(ctx, serviceId)
	if err != nil {
		return nil, err
	}
	if serviceDTO == nil {
		return nil, uds.ErrServiceNotFound
	}
	res := &UDSMapping{
		SID:         serviceDTO.SID,
		Name:        serviceDTO.Name,
		Description: serviceDTO.Description,
	}
	var subFunctionDTOs []SubfunctionDTO
	if serviceDTO.HasSubFunction {
		subFunctionDTOs, err = GetSubfunctionByServiceID(ctx, serviceDTO.SID)
		if err != nil {
			return nil, err
		}
		if len(subFunctionDTOs) != 0 {
			res.SubFunctions = subFunctionDTOs
		}
		for sf := range subFunctionDTOs {
			if subFunctionDTOs[sf].Value == message[1] {
				res.MatchedSubFunction = &subFunctionDTOs[sf]
				break
			}
		}
	}
	return res, nil
}
