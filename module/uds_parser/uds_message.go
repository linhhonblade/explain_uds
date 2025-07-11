package uds_parser

import (
	"database/sql"
	"fmt"
	"strings"
)

type Parameter struct {
	Type        string
	Value       string
	Description string
}

type UDSMapping struct {
	SID                string
	Name               string
	Description        sql.NullString
	SubFunctions       []SubfunctionDTO
	MatchedSubFunction *SubfunctionDTO
	Parameters         []Parameter // Sub-function from message (if any)
	MatchedParameters  []Parameter // Parameters matched from message (e.g., DIDs)
}

func (m *UDSMapping) String() string {
	if m == nil {
		return "nil"
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Service ID: %s\n", m.SID)
	fmt.Fprintf(&b, "Service Name: %s\n", m.Name)
	if m.Description.Valid {
		fmt.Fprintf(&b, "Service Description: %s\n", m.Description.String)
	}
	if m.SubFunctions != nil {
		fmt.Fprintf(&b, "Sub Functions:\n")
		for _, sf := range m.SubFunctions {
			fmt.Fprintf(&b, "  - %s (Value: %s, Description: %s)\n", sf.Name, sf.Value, sf.Description.String)
		}
	}
	if m.MatchedSubFunction != nil {
		fmt.Fprintf(&b, "Sub Function: %s (Value: %s, Description: %s)\n",
			m.MatchedSubFunction.Name, m.MatchedSubFunction.Value, m.MatchedSubFunction.Description.String)
	} else {
		fmt.Fprint(&b, "No matched sub function.\n")
	}
	return b.String()

}
