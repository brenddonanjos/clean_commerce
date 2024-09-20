package utils

import (
	"database/sql"
	"time"
)

func StringToSqlNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: s, Valid: false}
	}

	return sql.NullString{String: s, Valid: true}
}

func TimeToSqlNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{Time: t, Valid: false}
	}

	return sql.NullTime{Time: t, Valid: true}
}

func StringToTime(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}
	}

	return t
}
