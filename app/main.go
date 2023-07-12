package main

import(
	"github.com/gofiber/fiber/v2"
	"fiber-app/router"
)

func main(){

	app:= fiber.New()

	rootRouter.SetRoot(app)


	app.Listen(":3000")


}
