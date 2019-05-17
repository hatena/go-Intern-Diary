package model

type PageInfo struct {
	TotalPage       int
	CurrentPage     int
	HasNextPage     bool
	HasPreviousPage bool
}

const ARTICLE_PAGE_LIMIT = 5

type Pager struct {
	Page       int
	TotalCount int
	Limit      int
}

func (p *Pager) HasNextPage() bool {
	return p.Page*p.Limit < p.TotalCount
}
func (p *Pager) HasPreviousPage() bool {
	return p.Page > 1
}

func (p *Pager) TotalPage() int {
	if p.TotalCount == 0 {
		return 1
	}
	return (p.TotalCount-1)/p.Limit + 1
}
