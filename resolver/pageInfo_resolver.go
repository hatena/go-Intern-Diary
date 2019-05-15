package resolver

import (
	"context"

	"github.com/hatena/go-Intern-Diary/model"
)

type pageInfoResolver struct {
	pageInfo *model.PageInfo
}

func (p *pageInfoResolver) TotalPage(ctx context.Context) int32 {
	return int32(p.pageInfo.TotalPage)
}

func (p *pageInfoResolver) CurrentPage(ctx context.Context) int32 {
	return int32(p.pageInfo.CurrentPage)
}

func (p *pageInfoResolver) HasNextPage(ctx context.Context) bool {
	return p.pageInfo.HasNextPage
}

func (p *pageInfoResolver) HasPreviousPage(ctx context.Context) bool {
	return p.pageInfo.HasPreviousPage
}
