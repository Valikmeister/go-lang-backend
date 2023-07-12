package rootRouter

import (
	"github.com/gofiber/fiber/v2"
	"fiber-app/handler"
)

func SetRoot(app *fiber.App){

	root:=app.Group("/")

	root.Get("/hello", rootHandler.Hello)





}
