package repository

import (
	"time"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/jmoiron/sqlx"
)

const RECOMMEND_LINIT_COUNT = 3

func (r *repository) ListRecommendedDiaries(diaryID uint64) ([]*model.Diary, error) {
	tagsOfDiary, err := r.getTags(diaryID)
	if err != nil {
		return nil, err
	}
	diaries, err := r.unionCategory(tagsOfDiary, diaryID)
	if err != nil {
		return nil, err
	}
	return diaries, nil
}

func (r *repository) sameTagDiaries(tags []*model.Tag, diaryID uint64) ([]*model.Diary, error) {
	if len(tags) == 0 {
		return nil, nil
	}
	tagIDs := make([]uint64, len(tags))
	for i, tag := range tags {
		tagIDs[i] = tag.ID
	}
	query, args, err := sqlx.In(
		`SELECT DISTINCT diary.id, diary.name, diary.user_id, diary.updated_at FROM diary
			JOIN diary_tag ON diary.id = diary_tag.diary_id
			JOIN user ON diary.user_id = user.id
			WHERE tag_id IN (?) AND diary.id != ?
			ORDER BY diary.updated_at DESC LIMIT 3
		`, tagIDs, diaryID,
	)
	if err != nil {
		return nil, err
	}
	var diaries []*model.Diary
	err = r.db.Select(&diaries, query, args...)
	return diaries, err
}

func (r *repository) unionCategory(tags []*model.Tag, diaryID uint64) ([]*model.Diary, error) {
	if len(tags) == 0 {
		return nil, nil
	}
	categoryIDs := make([]int, len(tags))
	for i, tag := range tags {
		categoryIDs[i] = tag.CategoryID
	}
	query, args, err := sqlx.In(
		`SELECT COUNT(*) as count, diary.id, diary.name, diary.user_id, diary.updated_at FROM diary 
			JOIN diary_tag ON diary.id = diary_tag.diary_id 
			JOIN tag ON tag.id = diary_tag.tag_id 
			JOIN user ON diary.user_id = user.id 
			WHERE tag.category_id IN (?) 
				AND diary.id != ? 
			GROUP BY diary.id
			ORDER BY count DESC, updated_at LIMIT ?
		`, categoryIDs, diaryID, RECOMMEND_LINIT_COUNT,
	)
	if err != nil {
		return nil, err
	}

	var rec []*record
	err = r.db.Select(&rec, query, args...)
	if err != nil {
		return nil, err
	}
	diaries := getDiaries(rec)
	return diaries, err
}

type record struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	UserID    uint64    `db:"user_id"`
	UpdatedAt time.Time `db:"updated_at"`
	Count     int       `db:"count"`
}

func getDiaries(records []*record) []*model.Diary {
	diaries := make([]*model.Diary, len(records))
	for i, rec := range records {
		diaries[i] = &model.Diary{
			ID:        rec.ID,
			Name:      rec.Name,
			UserID:    rec.UserID,
			UpdatedAt: rec.UpdatedAt,
		}
	}
	return diaries
}
