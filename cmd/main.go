package main

import (
	// "github.com/diegolikescode/rinha-backend-q1-2024/domain/config"
	"fmt"

	"github.com/diegolikescode/rinha-backend-q1-2024/domain/external"
)

var ClientesArr []external.Cliente

func main() {
    // config.SetupCassandra()

    ClientesArr = append(ClientesArr, 
	external.Cliente{
	    ID: 1,
	    Nome: "baratao",
	    Limite: 1000 * 100,
	}, external.Cliente{
	    ID: 2,
	    Nome: "baratasso",
	    Limite: 100 * 100,
	}, external.Cliente{
	    ID: 3,
	    Nome: "baratissimo",
	    Limite: 100 * 100,
	}, external.Cliente{
	    ID: 4,
	    Nome: "baratoso",
	    Limite: 100 * 100,
	}, external.Cliente{
	    ID: 3,
	    Nome: "baratinho",
	    Limite: 100 * 100,
	},
    )

    fmt.Println(ClientesArr)

    external.SetupFiber()
}

