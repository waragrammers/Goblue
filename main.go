package main

import (
	_ "github.com/lib/pq"
	"github.com/waragrammers/Goblue/dbdriver"
)

func init() {
	dbdriver.PostgresDb()
}

func main() {

}
