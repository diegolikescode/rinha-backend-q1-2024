package external

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

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
    DataExtrato   int32 `json:"data_extrato"`
    Limite	  int32 `json:"limite"`
}

type Extrato struct {
    Saldo	        Saldo `json:"saldo"`
    UltimasTransacoes   []Transacao `json:"ultimas_transacoes"`
}

func HttpTransacoes(c fiber.Ctx) error {
    userID := c.Params("id")
    fmt.Println(userID)

    return c.Status(fiber.StatusOK).SendString("OKAY BEIBE ALRAITE")
}

