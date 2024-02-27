package external

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/diegolikescode/rinha-backend-q1-2024/domain/config"
	"github.com/gofiber/fiber/v3"
)

func NovaTransacao(c fiber.Ctx) error {
    log.Println("STARTING NovaTransacao")
    var parseErr error

    var t Transacao
    if parseErr = json.Unmarshal(c.Body(), &t); parseErr != nil {
	fmt.Println(parseErr)
	return c.SendStatus(fiber.StatusInternalServerError)
    }

    var userID int
    if userID, parseErr = strconv.Atoi(c.Params("id")); parseErr != nil {
	log.Fatal("o ID do cliente nao eh um integer valido", userID)
    }

    var conta Conta
    if t.Tipo == "c" {
	getIt, err := InserirCredito.Exec(userID, t.Valor, t.Descricao)
	if err != nil {
	    log.Fatal("not able to getIt")
	}
	fmt.Println(getIt)
    } else {
	getIt, err := InserirDebito.Exec(userID, t.Valor, t.Descricao)
	if err != nil {
	    log.Fatal("not able to getIt")
	}
	fmt.Println(getIt)
    }

    log.Println("entrypoint NovaTransacao executado com sucesso")
    return c.Status(fiber.StatusOK).JSON(conta)
}

func ClienteExtrato(c fiber.Ctx) error {
    var parseErr error
    var userID int
    if userID, parseErr = strconv.Atoi(c.Params("id")); parseErr != nil {
	log.Fatal("o ID do cliente nao eh um integer valido", userID)
    }

    if userID < 1 || userID > 5 {
	return c.SendStatus(fiber.StatusNotFound)
    }

    var wg sync.WaitGroup
    wg.Add(2)

    var extrato Extrato
    go buscaExtrato(&wg, &extrato, &userID)
    go buscaTransacoes(&wg, &extrato, &userID)
    wg.Wait()
    extrato.Saldo.DataExtrato = TimeNowFormatted()

    return c.Status(fiber.StatusOK).JSON(extrato)
}

func buscaExtrato(wg *sync.WaitGroup, extrato *Extrato, userID *int) {
    log.Println("inicia buscaExtrato")
    defer wg.Done()

    err := config.Session.QueryRow(
	"SELECT limite, saldo FROM clientes WHERE id = $1", userID).Scan(
	&extrato.Saldo.Limite, &extrato.Saldo.Total) 
    if err != nil {
	log.Fatal("ERROR QueryRow ", err)
    }

    log.Println("finaliza buscaExtrato")
}

func buscaTransacoes(wg *sync.WaitGroup, extrato *Extrato, userID *int) {
    log.Println("inicia buscaTransacoes")
    defer wg.Done()

    rows, err := config.Session.Query(
	`SELECT valor, tipo, descricao, realizada_em
	FROM transacoes
	WHERE id_cliente = $1
	ORDER BY realizada_em DESC
	LIMIT 10`, userID)
    if err != nil {
	log.Fatal("ERROR Postgres buscaTransacoes", err)
    }

    for rows.Next() {
	var t Transacao 
	if err := rows.Scan(&t.Valor, &t.Tipo, &t.Descricao, &t.RealizadaEm); err != nil {
	    log.Fatal("ERROR Scan buscarTransacoes")
	}
	extrato.UltimasTransacoes = append(extrato.UltimasTransacoes, t)
    }

    log.Println("finaliza buscaTransacoes")
}

func TimeNowFormatted() string {
    return time.Now().Format("2006-01-02T15:04:05.999999Z")
}

