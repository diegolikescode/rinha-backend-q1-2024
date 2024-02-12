package external

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
)


func HttpTransacoes(c fiber.Ctx) error {
    var t Transacao
    if err := json.Unmarshal(c.Body(), &t); err != nil {
	fmt.Println(err)
	return c.Status(fiber.StatusInternalServerError).SendString("WRONG!")
    }

    for _, cli := range(ClientesArr) {
    	
    }

    fmt.Println(t)

    return c.Status(fiber.StatusOK).JSON(t)
}
