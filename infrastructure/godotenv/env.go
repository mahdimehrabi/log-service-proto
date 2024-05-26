package godotenv

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Env struct {
	PGSQLConnection string
	Port            string
}

func (e *Env) Load() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	e.PGSQLConnection = os.Getenv("PGSQLCONNECTION")
	e.Port = os.Getenv("PORT")
}
