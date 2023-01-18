package datasource

import (
	"learning/package/datasource/sql"
)

type MongoDB struct{}

func NewMongoDB() {
	_ = sql.MySQLConfig{}
}
