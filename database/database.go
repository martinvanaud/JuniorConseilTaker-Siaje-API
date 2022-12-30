package database

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"juniorconseiltaker-siaje-api/database/entities"
)

type Database struct {
	Def *entities.Profile
}
