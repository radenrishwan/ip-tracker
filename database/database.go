package database

import (
	"fmt"
	"github.com/ip2location/ip2location-go/v9"
	"os"
)

func NewDatabase() *ip2location.DB {
	dir, err := os.Getwd()

	// change location to db path
	db, err := ip2location.OpenDB(fmt.Sprintf("%s/database/IP2LOCATION-LITE-DB11.IPV6.BIN", dir))

	if err != nil {
		panic(err)
	}

	return db
}
