package uds

import "errors"

var (
	ErrEmptyMessage           = errors.New("empty UDS message")
	ErrTooShortForSubFunction = errors.New("message too short for sub-function")
	ErrNegativeResponseShort  = errors.New("negative response too short")
	ErrServiceNotFound        = errors.New("service not found in database")
)
