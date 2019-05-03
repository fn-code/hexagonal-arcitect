package main

import (
	"context"
	"log"
	"os"

	"github.com/fn-code/hexagonal-arcitec/storage/psql"
)

func main() {
	pgconn, err := psql.NewPostgres(context.Background(), os.Getenv("PGUSERNAME"), os.Getenv("PGPASSWORD"), os.Getenv("PGPORT"), os.Getenv("PGDATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	defer pgconn.Close()
}
