package config

import "database/sql"

func NewInt32Optional(data sql.NullInt64) *int32 {
	var value *int32
	if data.Valid {
		typeInt32 := int32(data.Int64)
		value = &typeInt32
	}
	return value
}

func NewStringOptional(data sql.NullString) *string {
	var value *string
	if data.Valid {
		value = &data.String
	}
	return value
}
