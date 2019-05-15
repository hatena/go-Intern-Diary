package repository

import (
	"database/sql"
	"time"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/jmoiron/sqlx"
)

var articleNotFoundError = model.NotFoundError("article")

func (r *repository) ListArticlesByDiaryID(diaryID uint64, page, limit int) ([]*model.Article, *model.PageInfo, error) {
	offset := (page - 1) * limit
	total, err := r.getArticleTotalCount(diaryID)

	if err != nil {
		return nil, nil, err
	}
	articles := make([]*model.Article, 0, limit)
	err = r.db.Select(&articles,
		`SELECT id, title, content, diary_id, updated_at FROM article
			WHERE diary_id = ?
			ORDER BY updated_at DESC LIMIT ? OFFSET ?`,
		diaryID, limit, offset,
	)
	if err != nil {
		return nil, nil, err
	}
	pager := model.Pager{page, total, model.ARTICLE_PAGE_LIMIT}
	pageInfo := &model.PageInfo{
		TotalPage:       pager.TotalPage(),
		CurrentPage:     page,
		HasNextPage:     pager.HasNextPage(),
		HasPreviousPage: pager.HasPreviousPage(),
	}
	return articles, pageInfo, nil
}

func (r *repository) getArticleTotalCount(diaryID uint64) (int, error) {
	var count int
	err := r.db.Get(&count,
		`select count(*) as count from article where diary_id = ?`,
		diaryID,
	)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *repository) CreateNewArticle(diaryID uint64, title string, content string) (*model.Article, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO article 
			(id, title, content, diary_id, created_at, updated_at)
			VALUES(?,?,?,?,?,?)`,
		id, title, content, diaryID, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.Article{ID: id, Title: title, Content: content, DiaryID: diaryID, UpdatedAt: now}, nil
}

func (r *repository) FindArticleByID(articleID, diaryID uint64) (*model.Article, error) {
	var article model.Article
	err := r.db.Get(&article,
		`SELECT id, title, content, diary_id, updated_at FROM article
			WHERE id = ? AND diary_id = ? LIMIT 1`,
		articleID, diaryID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, articleNotFoundError
		}
		return nil, err
	}
	return &article, nil
}

func (r *repository) DeleteArticle(articleID uint64) (err error) {
	_, err = r.db.Exec(
		`DELETE FROM article
			WHERE id = ?`,
		articleID,
	)
	return
}

func (r *repository) ListArticlesByIDs(articleIDs []uint64) ([]*model.Article, error) {
	if len(articleIDs) == 0 {
		return nil, nil
	}
	articles := make([]*model.Article, len(articleIDs))
	query, args, err := sqlx.In(
		`SELECT id, title, content, diary_id, updated_at FROM article
			WHERE id IN (?)
			ORDER BY created_at DESC`, articleIDs,
	)
	if err != nil {
		return nil, err
	}
	err = r.db.Select(&articles, query, args...)
	return articles, err
}

func (r *repository) ListArticlesByDiaryIDs(diaryIDs []uint64) (map[uint64][]*model.Article, error) {
	if len(diaryIDs) == 0 {
		return nil, nil
	}
	query, args, err := sqlx.In(
		`SELECT id, title, content, diary_id, updated_at FROM article
			WHERE diary_id IN (?)`,
		diaryIDs,
	)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	articles := make(map[uint64][]*model.Article)
	for rows.Next() {
		var article model.Article
		rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.DiaryID,
			&article.UpdatedAt,
		)
		articles[article.DiaryID] = append(articles[article.DiaryID], &article)
	}
	return articles, nil
}

func (r *repository) UpdateArticle(articleID uint64, title, content string) (*model.Article, error) {
	var article model.Article
	err := r.db.Get(&article,
		`SELECT id, title, content, diary_id, updated_at FROM article
			WHERE id = ? LIMIT 1`,
		articleID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, articleNotFoundError
		}
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`UPDATE article SET title = ?, content = ?, updated_at = ?
			WHERE id = ?`,
		title, content, now, articleID,
	)
	if err != nil {
		return nil, err
	}
	return &model.Article{ID: article.ID, DiaryID: article.DiaryID, Title: title, Content: content, UpdatedAt: now}, nil
}
