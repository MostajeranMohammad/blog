package repo

import (
	"context"

	"github.com/MostajeranMohammad/blog/internal/entity"
	"gorm.io/gorm/clause"
)

type (
	BlogPost interface {
		Create(ctx context.Context, p map[string]interface{}) (entity.BlogPost, error)
		GetById(ctx context.Context, id int) (entity.BlogPost, error)
		GetByTitle(ctx context.Context, title string) (entity.BlogPost, error)
		GetAll(ctx context.Context, filter clause.AndConditions, skip int, limit int) ([]entity.BlogPost, error)
		Update(ctx context.Context, id int, m map[string]interface{}) (entity.BlogPost, error)
		Delete(ctx context.Context, id int) (entity.BlogPost, error)
	}
)
