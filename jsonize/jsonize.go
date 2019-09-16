// Package Jsonize JSON-izes sql.NullString and sql.NullInt into true values
package jsonize

import (
	"database/sql"
	"encoding/json"
)

type JsonNullInt64 struct {
	sql.NullInt64
}
type JsonNullString struct {
	sql.NullString
}
type JsonNullFloat64 struct {
	sql.NullFloat64
}

func (v JsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v JsonNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v JsonNullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	} else {
		return json.Marshal(nil)
	}
}
