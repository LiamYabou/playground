package main

import (
	"fmt"

	"github.com/LiamYabou/playground/go/module/pkg/db"

	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := db.Open()
	err = dbConn.Ping()
	if err != nil {
		fmt.Errorf("Could not open a connection, err: %s", err)
	} else {
		fmt.Println("Connect the DB succefully!")
	}
}
