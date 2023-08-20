module brews

go 1.21.0

replace tobrew/types/server => ../../types/server

replace tobrew/types/tobrew => ../../types/tobrew

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/google/uuid v1.3.0
	tobrew/types/server v0.0.0-00010101000000-000000000000
)

require (
	github.com/jmoiron/sqlx v1.3.5 // indirect
	tobrew/types/tobrew v0.0.0-00010101000000-000000000000 // indirect
)
