package uds_parser

import "database/sql"

type ServiceDTO struct {
	SID              string
	Name             string
	HasSubFunction   bool
	PositiveResponse sql.NullInt64
	Description      sql.NullString
}

type SubfunctionDTO struct {
	SID         string
	Value       string
	Name        string
	Description sql.NullString
}
