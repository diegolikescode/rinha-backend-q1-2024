package external

import (
	"database/sql"
	"log"

	"github.com/diegolikescode/rinha-backend-q1-2024/domain/config"
)

var (
    InserirCredito *sql.Stmt
    InserirDebito *sql.Stmt
    SelectUltimasTransacoes *sql.Stmt
)

type Cliente struct {
    ID      int32 `json:"id" db:"id"`
    Nome    string `json:"nome" db:"nome"`
    Limite  int32 `json:"limite" db:"limite"`
}

type Transacao struct {
    Valor      int32 `json:"valor"`
    Tipo       string `json:"tipo"`
    Descricao  string `json:"descricao"`
    RealizadaEm string `json:"realizada_em"`
}

type Conta struct {
    Limite   int32 `json:"limite"`
    Saldo    int32 `json:"saldo"`
}

type Saldo struct {
    Total         int32 `json:"total"`
    DataExtrato   string `json:"data_extrato"`
    Limite	  int32 `json:"limite"`
}

type Extrato struct {
    Saldo	        Saldo `json:"saldo"`
    UltimasTransacoes   []Transacao `json:"ultimas_transacoes"`
}

func DeclareStmts() {
    var err error
    InserirCredito, err = config.Session.Prepare(
	`SELECT * FROM inserir_credito($1, $2, $3)`)   
    if err != nil {
	log.Fatal("ERROR: InserirCredito ", err)
    }

    InserirDebito, err = config.Session.Prepare(
	`SELECT * FROM inserir_debito($1, $2, $3)`)   
    if err != nil {
	log.Fatal("ERROR: inserir_debito ", err)
    }

    SelectUltimasTransacoes, err = config.Session.Prepare(
	`SELECT * FROM obter_ultimas_transacoes($1)`)
    if err != nil {
	log.Fatal("ERROR: obter_ultimas_transacoes ", err)
    }

}

