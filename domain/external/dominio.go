package external

import (
	"database/sql"
	"log"

	"github.com/diegolikescode/rinha-backend-q1-2024/domain/config"
	"github.com/go-playground/validator/v10"
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
    Valor      int32 `json:"valor" validate:"required,fieldexcludes=., "`
    Tipo       string `json:"tipo" validate:"required,oneof=c d"`
    Descricao  string `json:"descricao" validate:"required,containsany,len=10"`
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

type FieldValidator struct {
    validator *validator.Validate
}

var validate = validator.New()

func (v FieldValidator) IsInputValid(data interface{}) bool {
    errs := validate.Struct(data)
    if errs != nil {
	return false
    }

    return true
}

func DeclareStmts() {
    var err error
    InserirCredito, err = config.Session.Prepare(
	`SELECT * FROM inserir_credito($1, $2, $3)`)   
    if err != nil {
	log.Println("ERROR: inserir_credito ", err)
    }

    InserirDebito, err = config.Session.Prepare(
	`SELECT * FROM inserir_debito($1, $2, $3)`)   
    if err != nil {
	log.Println("ERROR: inserir_debito ", err)
    }

    SelectUltimasTransacoes, err = config.Session.Prepare(
	`SELECT * FROM obter_ultimas_transacoes($1)`)
    if err != nil {
	log.Println("ERROR: obter_ultimas_transacoes ", err)
    }

}

