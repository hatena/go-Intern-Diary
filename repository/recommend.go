package repository

import (
	"github.com/hatena/go-Intern-Diary/model"
	"github.com/jmoiron/sqlx"
)

func (r *repository) ListRecommendedDiaries(diaryID uint64) ([]*model.Diary, error) {
	tagsOfDiary, err := r.getTags(diaryID)
	if err != nil {
		return nil, err
	}
	diaries, err := r.sameTagDiaries(tagsOfDiary, diaryID)
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
	categoryIDs := make([]uint64, len(tags))
	for i, tag := range tags {
		categoryIDs[i] = tag.ID
	}
	query, args, err := sqlx.In(
		`SELECT DISTINCT diary.id, diary.name, diary.user_id, diary.updated_at FROM diary
			JOIN diary_tag ON diary.id = diary_tag.diary_id
			JOIN tag ON tag.id = diary_tag.tag_id
			JOIN user ON diary.user_id = user.id
			WHERE tag.category_id IN (?) AND diary.id != ?
		`, categoryIDs, diaryID,
	)
	if err != nil {
		return nil, err
	}
	var diaries []*model.Diary
	err = r.db.Select(&diaries, query, args...)
	return diaries, err
}
