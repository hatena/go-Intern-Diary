package resolver

import (
	"context"

	"github.com/hatena/go-Intern-Diary/model"
)

type categoryResolver struct {
	category *model.Category
}

func (c *categoryResolver) ID(ctx context.Context) int32 {
	return int32(c.category.ID)
}

func (c *categoryResolver) CategoryName(ctx context.Context) string {
	return c.category.CategoryName
}
