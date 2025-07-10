package uds_parser

import "database/sql"

type SubFunction struct {
	Value       string
	Name        string
	Description string
}

type Parameter struct {
	Type        string
	Value       string
	Description string
}
type UDSMapping struct {
	SID                string
	Name               string
	Description        sql.NullString
	SubFunctions       []SubFunction
	MatchedSubFunction *SubFunction
	Parameters         []Parameter // Sub-function from message (if any)
	MatchedParameters  []Parameter // Parameters matched from message (e.g., DIDs)
}
