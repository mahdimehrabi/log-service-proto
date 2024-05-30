package godotenv

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Env struct {
	PGSQLConnection string
	ServerAddr      string
}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) Load() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	e.PGSQLConnection = os.Getenv("PGSQLCONNECTION")
	e.ServerAddr = os.Getenv("ServerAddr")
}
