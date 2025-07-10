package uds_parser

import (
	"context"
	"database/sql"
	"errors"
	uds "explain_uds/common"
	"log"
)

func getDB(ctx context.Context) *sql.DB {
	db, ok := ctx.Value(uds.CtxKeyDB{}).(*sql.DB)
	if !ok {
		log.Fatalln("no *sql.DB in context")
	}
	return db
}

func GetServiceByID(ctx context.Context, serviceID string) (*ServiceDTO, error) {
	db := getDB(ctx)
	var dto ServiceDTO

	// To prevent Null value in DB is restured as 0
	// sqlite only have INTEGER 64bit
	// backward compatible with older go version where database/sql does not support NullInt32 and NullInt16

	row := db.QueryRow(`select sid, name, has_subfunction, positive_response, description from services where sid = ?`, serviceID)
	err := row.Scan(&dto.SID, &dto.Name, &dto.HasSubFunction, &dto.PositiveResponse, &dto.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// If no rows are found, return nil and no error
			return nil, nil
		}
		return nil, err
	}
	return &dto, nil
}

func GetSubfunctionByServiceID(ctx context.Context, serviceID string) ([]SubfunctionDTO, error) {
	db := getDB(ctx)

	rows, err := db.Query(`select value, name, sid, description from sub_functions where sid = ?`, serviceID)
	if err != nil {
		return nil, err
	}
	var res []SubfunctionDTO
	for rows.Next() {
		var dto SubfunctionDTO
		if err := rows.Scan(&dto.Value, &dto.Name, &dto.SID, &dto.Description); err != nil {
			return nil, err
		}
		res = append(res, dto)
	}

	// Check if the loop exited due to an error or because there are no more rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
