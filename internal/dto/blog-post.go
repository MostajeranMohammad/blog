package dto

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

type (
	CreateNewBlogPostDto struct {
		Title    string `validate:"required"`
		AuthorId uint   `validate:"required"`
		Content  string `validate:"required"`
	}
	UpdateBlogPostDto struct {
		Title    string
		AuthorId uint `validate:"gt=0"`
		Content  string
	}
	FilterBlogPosts struct {
		AuthorId uint   `query:"author_id" validate:"gt=0"`
		FromDate string `query:"from_date"`
		ToDate   string `query:"to_date"`
	}
)

func (q FilterBlogPosts) ToQueryModel() (clause.AndConditions, error) {
	queryModel := clause.AndConditions{}

	if q.AuthorId > 0 {
		queryModel.Exprs = append(queryModel.Exprs, clause.Eq{Column: "author_id", Value: q.AuthorId})
	}
	if q.FromDate != "" {
		t, err := time.Parse("2006-01-02T15:04:05.000Z", q.FromDate)
		if err != nil {
			return queryModel, fiber.NewError(400, "invalid date format")
		}

		queryModel.Exprs = append(queryModel.Exprs, clause.Gte{Column: "created_at", Value: t})
	}
	if q.ToDate != "" {
		t, err := time.Parse("2006-01-02T15:04:05.000Z", q.ToDate)
		if err != nil {
			return queryModel, fiber.NewError(400, "invalid date format")
		}

		queryModel.Exprs = append(queryModel.Exprs, clause.Lte{Column: "created_at", Value: t})
	}

	return queryModel, nil
}
