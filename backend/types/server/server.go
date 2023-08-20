package server

import "github.com/jmoiron/sqlx"

type Server struct {
	Db *sqlx.DB
}
