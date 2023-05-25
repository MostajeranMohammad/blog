package routes

import (
	"github.com/MostajeranMohammad/blog/internal/controller"
	"github.com/MostajeranMohammad/blog/pkg/guards"
	"github.com/gofiber/fiber/v2"
)

func NewStaticFileRouter(controller controller.BlogPost, jwtGuard guards.JWT) *fiber.App {
	router := fiber.New()

	router.Post("", jwtGuard.GetStrictJWTGuard(), controller.CreateNewBlogPost)
	router.Get("", jwtGuard.GetOptionalJWTGuard(), controller.GetAllBlogPosts)
	router.Get("/:id", jwtGuard.GetOptionalJWTGuard(), controller.GetBlogPostById)
	router.Get("/get-by-title/:title", jwtGuard.GetOptionalJWTGuard(), controller.GetBlogPostByName)
	router.Put("/:id", jwtGuard.GetStrictJWTGuard(), controller.UpdateBlogPost)
	router.Delete("/:id", jwtGuard.GetStrictJWTGuard(), controller.DeletePost)

	return router
}
