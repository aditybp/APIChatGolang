package routes

import (
	"APICHATGOLANG/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(App *fiber.App) {

	App.Post("/api/Login", controllers.LoginAuth)
	App.Post("/api/register", controllers.Register)
	App.Get("/api/user", controllers.GetUser)
	App.Get("/api/getmessages", controllers.GetAllMessages)
	App.Post("/api/postmessages", controllers.MakeMessages)

}
