package types

import "github.com/jmoiron/sqlx"

type Server struct {
	Db *sqlx.DB
}
