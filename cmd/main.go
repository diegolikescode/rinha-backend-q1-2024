package main

import (
	"github.com/diegolikescode/rinha-backend-q1-2024/domain/config"
	"github.com/diegolikescode/rinha-backend-q1-2024/domain/external"
)

func main() {
    config.SetupPostgres()
    external.DeclareStmts()
    external.SetupFiber()
}

