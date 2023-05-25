package controller

import "github.com/gofiber/fiber/v2"

type (
	BlogPost interface {
		CreateNewBlogPost(c *fiber.Ctx) error
		GetBlogPostById(c *fiber.Ctx) error
		GetBlogPostByName(c *fiber.Ctx) error
		GetAllBlogPosts(c *fiber.Ctx) error
		UpdateBlogPost(c *fiber.Ctx) error
		DeletePost(c *fiber.Ctx) error
	}
)
