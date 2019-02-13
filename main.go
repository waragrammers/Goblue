package main

import (
	_ "github.com/lib/pq"
	"github.com/waragrammers/Goblue/dbdriver"
)

func main() {

	dbdriver.PostgresDb()

}
