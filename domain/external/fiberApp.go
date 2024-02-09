package external

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func SetupFiber() {
    app := fiber.New()
    app.Get("/clientes/:id/transacoes", HttpTransacoes)

    log.Fatal(app.Listen(":6969"))
}
