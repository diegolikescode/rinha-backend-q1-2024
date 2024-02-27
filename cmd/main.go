package main

import (
	"fmt"

	"github.com/diegolikescode/rinha-backend-q1-2024/domain/config"
	"github.com/diegolikescode/rinha-backend-q1-2024/domain/external"
)

func main() {
    config.SetupPostgres()
    fmt.Println("SetupPostgres")
    external.DeclareStmts()
    fmt.Println("DeclareStmts")
    external.SetupFiber()
    fmt.Println("SetupFiber")
}

