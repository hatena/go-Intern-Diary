package repository

import "github.com/hatena/go-Intern-Diary/model"

func (r *repository) ListArticlesByDiaryID(diaryID, limit, offset uint64) ([]*model.Article, error) {
	articles := make([]*model.Article, 0, limit)
	err := r.db.Select(&articles,
		`SELECT id, title, content, diary_id, updated_at FROM article
			WHERE diary_id = ?
			ORDER BY updated_at DESC LIMIT ? OFFSET ?`,
		diaryID, limit, offset,
	)
	return articles, err
}
