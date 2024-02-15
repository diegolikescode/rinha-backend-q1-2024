package external

import (
	"database/sql"
	"log"

	"github.com/diegolikescode/rinha-backend-q1-2024/domain/config"
)

var (
    InsertStmt *sql.Stmt
    UpdateStmt *sql.Stmt
    SelectTransacoesStmt *sql.Stmt
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
    InsertStmt, err = config.Session.Prepare(
	`INSERT INTO transacoes (id_cliente, valor, tipo, descricao, realizada_em)
	VALUES ($1, $2, $3, $4, $5)`)   
    if err != nil {
	log.Fatal("ERROR: insertStmt ", err)
    }

    UpdateStmt, err = config.Session.Prepare(`
	UPDATE clientes 
	SET saldo = $1
	WHERE id = $2;`)
    if err != nil {
	log.Fatal("ERROR: insertStmt ", err)
    }

}

