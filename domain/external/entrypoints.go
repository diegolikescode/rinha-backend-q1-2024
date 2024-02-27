package external

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diegolikescode/rinha-backend-q1-2024/domain/config"
	"github.com/gofiber/fiber/v3"
)

var LocalValidator FieldValidator

func NovaTransacao(c fiber.Ctx) error {
    log.Println("STARTING NovaTransacao")
    var parseErr error


    var t Transacao
    if parseErr = json.Unmarshal(c.Body(), &t); parseErr != nil {
	fmt.Println(parseErr)
	return c.SendStatus(fiber.StatusInternalServerError)
    }

    if(!LocalValidator.IsInputValid(t)) {
	return c.SendStatus(fiber.StatusUnprocessableEntity)
    }

    var userID int
    if userID, parseErr = strconv.Atoi(c.Params("id")); parseErr != nil {
	log.Println("o ID do cliente nao eh um integer valido", userID)
    }

    var conta Conta
    if t.Tipo == "c" {
	err := InserirCredito.QueryRow(userID, t.Valor, t.Descricao).Scan(&conta.Saldo, &conta.Limite)
	if err != nil {
	    log.Println("ERROR: InserirCredito:: ", err)
	    if strings.Contains(err.Error(), "NOUSER") {
		return c.SendStatus(fiber.StatusNotFound)
	    }
	}
	InserirCredito.Close()

    } else {
	err := InserirDebito.QueryRow(userID, t.Valor, t.Descricao).Scan(&conta.Saldo, &conta.Limite)
	if err != nil {
	    log.Println("ERROR: InserirDebito:: ", err)
	    if strings.Contains(err.Error(), "NOUSER") {
		return c.SendStatus(fiber.StatusNotFound)
	    }

	    if strings.Contains(err.Error(), "NOLIMIT") {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	    }
	}
	InserirDebito.Close()
    }

    log.Println("entrypoint NovaTransacao executado com sucesso")
    return c.Status(fiber.StatusOK).JSON(conta)
}

func ClienteExtrato(c fiber.Ctx) error {
    var parseErr error
    var userID int
    if userID, parseErr = strconv.Atoi(c.Params("id")); parseErr != nil {
	log.Println("o ID do cliente nao eh um integer valido", userID)
    }

    rows, err := SelectUltimasTransacoes.Query(userID)
    if err != nil {
	log.Println("ERROR: SelectUltimasTransacoes:: ", err)
	if strings.Contains(err.Error(), "NOUSER") {
	    return c.SendStatus(fiber.StatusNotFound)
	}
    }

    var extrato Extrato
    for rows.Next() {
	var t Transacao
	rows.Scan(&t.Valor, &t.Tipo, &t.Descricao, &t.RealizadaEm, &extrato.Saldo.Total, &extrato.Saldo.Limite)
	extrato.UltimasTransacoes = append(extrato.UltimasTransacoes, t)
    }

    rows.Close()
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
	log.Println("ERROR QueryRow ", err)
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
	log.Println("ERROR Postgres buscaTransacoes", err)
    }

    for rows.Next() {
	var t Transacao 
	if err := rows.Scan(&t.Valor, &t.Tipo, &t.Descricao, &t.RealizadaEm); err != nil {
	    log.Println("ERROR Scan buscarTransacoes")
	}
	extrato.UltimasTransacoes = append(extrato.UltimasTransacoes, t)
    }

    log.Println("finaliza buscaTransacoes")
}

func TimeNowFormatted() string {
    return time.Now().Format("2006-01-02T15:04:05.999999Z")
}

