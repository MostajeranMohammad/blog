package repo

import (
	"github.com/MostajeranMohammad/blog/internal/entity"
	"gorm.io/gorm/clause"
)

type (
	BlogPost interface {
		Create(p entity.BlogPost) (entity.BlogPost, error)
		GetById(id int) (entity.BlogPost, error)
		GetByTitle(title string) (entity.BlogPost, error)
		GetAll(filter clause.AndConditions, skip int, limit int) ([]entity.BlogPost, error)
		Update(id int, m entity.BlogPost) (entity.BlogPost, error)
		Delete(id int) (entity.BlogPost, error)
	}
)
