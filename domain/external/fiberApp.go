package external

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func SetupFiber() {
    app := fiber.New()
    app.Post("/clientes/:id/transacoes", NovaTransacao)
    app.Get("/clientes/:id/extrato", ClienteExtrato)

    log.Fatal(app.Listen(":6969"))
}

