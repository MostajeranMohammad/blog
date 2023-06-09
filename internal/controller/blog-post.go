package controller

import (
	"encoding/json"

	"github.com/MostajeranMohammad/blog/internal/dto"
	"github.com/MostajeranMohammad/blog/internal/entity"
	"github.com/MostajeranMohammad/blog/internal/usecase"
	"github.com/MostajeranMohammad/blog/pkg/logger"
	"github.com/MostajeranMohammad/blog/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type BlogPostController struct {
	useCase usecase.BlogPost
	logger  logger.Logger
}

func NewBlogPostController(useCase usecase.BlogPost, logger logger.Logger) BlogPost {
	return &BlogPostController{useCase, logger}
}

// @Accept       json
// @Produce      json
// @Param        body  body  dto.CreateNewBlogPostDto  true "no comment"
// @Router       /blog-post [post]
// @Security BearerAuth
func (bc *BlogPostController) CreateNewBlogPost(c *fiber.Ctx) error {
	body := dto.CreateNewBlogPostDto{}
	c.BodyParser(&body)
	err := utils.ValidateDto(body)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	mapBody := make(map[string]interface{})
	err = json.Unmarshal(c.Body(), &mapBody)
	if err != nil {
		return fiber.NewError(400, "failed to unmarshal json data")
	}

	result, err := bc.useCase.CreateNewBlogPost(c.Context(), mapBody)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: result})
}

// @Produce      json
// @Param        id  path  int  true  "no comment"
// @Router       /blog-post/{id} [get]
func (bc *BlogPostController) GetBlogPostById(c *fiber.Ctx) error {
	postId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	result, err := bc.useCase.GetBlogPostById(c.Context(), postId)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: result})
}

// @Produce      json
// @Param        title  path  string  true "no comment"
// @Router       /blog-post/get-by-title/{title} [get]
func (bc *BlogPostController) GetBlogPostByName(c *fiber.Ctx) error {
	postTitle := c.Params("title")

	result, err := bc.useCase.GetBlogPostByTitle(c.Context(), postTitle)
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
func (bc *BlogPostController) GetAllBlogPosts(c *fiber.Ctx) error {
	queries := dto.FilterBlogPosts{}
	err := c.QueryParser(&queries)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	limit, skip := c.QueryInt("limit"), c.QueryInt("skip")

	result, err := bc.useCase.GetAllBlogPosts(c.Context(), queries, limit, skip)
	if err != nil {
		return err
	}
	count, err := bc.useCase.Count(c.Context(), queries)
	if err != nil {
		bc.logger.Error(err.Error())
	}

	return c.JSON(entity.ResponseModel{Successful: true, Data: result, Meta: map[string]interface{}{
		"returnedCount": len(result),
		"count":         count,
	}})
}

// @Accept       json
// @Produce      json
// @Param        body  body  dto.UpdateBlogPostDto  true "no comment"
// @Param        id path  int  true  "no comment"
// @Router       /blog-post/{id} [put]
// @Security BearerAuth
func (bc *BlogPostController) UpdateBlogPost(c *fiber.Ctx) error {
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

	mapBody := make(map[string]interface{})
	err = json.Unmarshal(c.Body(), &mapBody)
	if err != nil {
		return fiber.NewError(400, "failed to unmarshal json data")
	}

	result, err := bc.useCase.UpdateBlogPost(c.Context(), id, mapBody)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: result})
}

// @Produce      json
// @Param        id path  int  true "no comment"
// @Router       /blog-post/{id} [delete]
// @Security BearerAuth
func (bc *BlogPostController) DeletePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	result, err := bc.useCase.DeletePost(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(entity.ResponseModel{Successful: true, Data: result})
}
