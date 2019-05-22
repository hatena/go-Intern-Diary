package repository

import (
	"github.com/hatena/go-Intern-Diary/model"
	"github.com/jmoiron/sqlx"
)

func (r *repository) ListTagsByIDs(tagIDs []uint64) ([]*model.Tag, error) {
	if len(tagIDs) == 0 {
		return nil, nil
	}
	tags := make([]*model.Tag, 0, len(tagIDs))
	query, args, err := sqlx.In(
		`SELECT id, tag_name, category_id FROM tag
			WHERE id IN (?)
			ORDER BY tag_name ASC`, tagIDs,
	)
	if err != nil {
		return nil, err
	}
	err = r.db.Select(&tags, query, args...)
	return tags, err
}
func (r *repository) ListTagsByDiaryIDs(diaryIDs []uint64) (map[uint64][]*model.Tag, error) {
	query, args, err := sqlx.In(
		`SELECT tag.id, tag_name, category_id, diary_id FROM tag
			JOIN diary_tag ON diary_tag.tag_id = tag.id
			WHERE diary_id IN (?)
			ORDER BY tag_name ASC
		`, diaryIDs,
	)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tags := make(map[uint64][]*model.Tag)
	for rows.Next() {
		var tag model.Tag
		var diaryID uint64
		rows.Scan(&tag.ID, &tag.TagName, &tag.CategoryID, &diaryID)
		tags[diaryID] = append(tags[diaryID], &tag)
	}
	return tags, nil
}
