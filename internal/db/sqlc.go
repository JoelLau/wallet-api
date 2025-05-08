package db

import dbgen "bank-app/internal/db/gen"

func NewSQLCRepository(dbtx dbgen.DBTX) *dbgen.Queries {
	return dbgen.New(dbtx)
}
