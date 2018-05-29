package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lab/ksp-api/router"
)

func main() {
	router.Routing()
}
