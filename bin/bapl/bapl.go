package main

import (
	bapl "bapl/lib"
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	var Here string=bapl.GetHereDirExe()

    var app *fiber.App=fiber.New(fiber.Config{
        CaseSensitive: true,
        ErrorHandler: func(c fiber.Ctx, err error) error {
            fmt.Println("fiber error")
            fmt.Println(err)
            return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
        },
    })

    app.Get("/*",static.New(filepath.Join(Here,"bapl-web/build")))

    app.Listen(":4206")
}