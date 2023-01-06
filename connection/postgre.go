package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnection(){
	var err error
	databaseUrl :="postgres://postgres:123@localhost:5432/personal-web"
	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to Database: %v", err)
		os.Exit(1)
	}
	fmt.Println("Database connected.")
}