module tobrew

go 1.21.0

replace tobrew/types/server => ./types/server

replace tobrew/types/tobrew => ./types/tobrew

replace tobrew/controllers/brews => ./controllers/brews

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-sql-driver/mysql v1.7.1
	github.com/jmoiron/sqlx v1.3.5
	github.com/joho/godotenv v1.5.1
	tobrew/controllers/brews v0.0.0-00010101000000-000000000000
	tobrew/types/server v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.3.0 // indirect
	tobrew/types/tobrew v0.0.0-00010101000000-000000000000 // indirect
)
