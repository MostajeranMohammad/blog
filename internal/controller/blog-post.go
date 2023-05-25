package controller

import (
	"github.com/MostajeranMohammad/blog/internal/dto"
	"github.com/MostajeranMohammad/blog/internal/entity"
	"github.com/MostajeranMohammad/blog/internal/usecase"
	"github.com/MostajeranMohammad/blog/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type BlogPostController struct {
	useCase usecase.BlogPost
}

func NewBlogPostController(useCase usecase.BlogPost) BlogPost {
	return &BlogPostController{useCase}
}

// @Accept       json
// @Produce      json
// @Param        body  body  dto.CreateNewBlogPostDto  true "no comment"
// @Router       /blog-post [post]
// @Security BearerAuth
func (bc BlogPostController) CreateNewBlogPost(c *fiber.Ctx) error {
	body := dto.CreateNewBlogPostDto{}
	c.BodyParser(&body)
	err := utils.ValidateDto(body)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	result, err := bc.useCase.CreateNewBlogPost(body)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: result})
}

// @Produce      json
// @Param        id  path  int  true  "no comment"
// @Router       /blog-post/{id} [get]
func (bc BlogPostController) GetBlogPostById(c *fiber.Ctx) error {
	postId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	result, err := bc.useCase.GetBlogPostById(postId)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: result})
}

// @Produce      json
// @Param        title  path  string  true "no comment"
// @Router       /blog-post/get-by-title/{title} [get]
func (bc BlogPostController) GetBlogPostByName(c *fiber.Ctx) error {
	postTitle := c.Params("title")

	result, err := bc.useCase.GetBlogPostByTitle(postTitle)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: result})
}

// @Produce      json
// @Param        limit  query  int  false "no comment"
// @Param        skip  query  int  false "no comment"
// @Param        author_id  query  int  false "no comment"
// @Param        from_date  query  string  false "no comment"
// @Param        to_date  query  string  false "no comment"
// @Router       /blog-post [get]
func (bc BlogPostController) GetAllBlogPosts(c *fiber.Ctx) error {
	queries := dto.FilterBlogPosts{}
	err := c.QueryParser(&queries)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	limit, skip := c.QueryInt("limit"), c.QueryInt("skip")

	result, err := bc.useCase.GetAllBlogPosts(queries, limit, skip)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: result})
}

// @Accept       json
// @Produce      json
// @Param        body  body  dto.UpdateBlogPostDto  true "no comment"
// @Param        id path  int  true  "no comment"
// @Router       /blog-post/{id} [put]
// @Security BearerAuth
func (bc BlogPostController) UpdateBlogPost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	body := dto.UpdateBlogPostDto{}
	c.BodyParser(&body)
	err = utils.ValidateDto(body)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	result, err := bc.useCase.UpdateBlogPost(id, body)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: result})
}

// @Produce      json
// @Param        id path  int  true "no comment"
// @Router       /blog-post/{id} [delete]
// @Security BearerAuth
func (bc BlogPostController) DeletePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	result, err := bc.useCase.DeletePost(id)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: result})
}
