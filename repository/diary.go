package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/hatena/go-Intern-Diary/instance"
	"github.com/hatena/go-Intern-Diary/model"
	"github.com/jmoiron/sqlx"
)

var diaryNotFoundError = model.NotFoundError("diary")

const DEFAULT_CATEGORY_ID = uint64(0)

func (r *repository) CreateNewDiary(userID uint64, name string, tagWithCategories []*model.TagWithCategory) (*model.Diary, error) { // Todo transaction and rollback
	now := time.Now()
	newDiary, err := r.insertDiary(userID, name, now)
	if err != nil {
		return nil, err
	}
	err = r.insertDiaryTags(newDiary.ID, tagWithCategories, now)
	if err != nil {
		return nil, err
	}
	return newDiary, nil
}

func (r *repository) insertDiary(userID uint64, name string, now time.Time) (*model.Diary, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(
		`INSERT INTO diary
			(id, name, user_id, created_at, updated_at)
			VALUES(?, ?, ?, ?, ?)
			`, id, name, userID, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.Diary{ID: id, Name: name, UserID: userID, UpdatedAt: now}, nil
}

func (r *repository) insertDiaryTags(diaryId uint64, tagWithCategories []*model.TagWithCategory, now time.Time) error {
	for _, tagWithCategory := range tagWithCategories {
		tagId, err := r.generateID()
		if err != nil {
			return err
		}
		if len(tagWithCategory.CategoryIDs) > 0 {
			_, err = r.db.Exec(
				`INSERT INTO tag
					(id, tag_name, category_id, created_at, updated_at)
					VALUES
					(?, ?, ?, ?, ?)
				`, tagId, tagWithCategory.TagName, tagWithCategory.CategoryIDs[0], now, now,
			)
			if err != nil {
				return err
			}
		} else {
			_, err = r.db.Exec(
				`INSERT INTO tag
					(id, tag_name, created_at, updated_at)
					VALUES
					(?, ?, ?, ?)
				`, tagId, tagWithCategory.TagName, now, now,
			)
			if err != nil {
				return err
			}
		}
		id, err := r.generateID()
		if err != nil {
			return err
		}
		_, err = r.db.Exec(
			`INSERT INTO diary_tag
			(id, diary_id, tag_id, updated_at, created_at)
			VALUES
			(?, ?, ?, ?, ?)
		`, id, diaryId, tagId, now, now,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) insertNewTags(newTagNames []string, now time.Time) ([]uint64, error) {
	ids := make([]uint64, len(newTagNames))
	for i, newTagName := range newTagNames {
		id, err := r.generateID()
		if err != nil {
			return []uint64{}, err
		}
		_, err = r.db.Exec(
			`INSERT INTO tag
			(id, tag_name, category_id, updated_at, created_at)
			VALUES
			(?, ?, ?, ?, ?)
		`, id, newTagName, DEFAULT_CATEGORY_ID, now, now,
		)
		if err != nil {
			return []uint64{}, err
		}
		ids[i] = id
	}
	return ids, nil
}

func (r *repository) newTagFilter(tags []string, storedTagMap map[string]uint64) []string {
	newTags := make([]string, 0, len(tags))
	for _, tagName := range tags {
		if _, ok := storedTagMap[tagName]; !ok {
			newTags = append(newTags, tagName)
		}
	}
	return newTags
}

func (r *repository) getTagIDsByNames(tagNames []string) (map[string]uint64, error) {
	type tagInfo struct {
		id      uint64
		tagName string
	}

	if len(tagNames) == 0 {
		return nil, nil
	}
	query, args, err := sqlx.In(
		`SELECT id, tag_name FROM tag
		WHERE tag_name IN (?)
		`,
		tagNames,
	)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tags := make(map[string]uint64)
	for rows.Next() {
		var tag tagInfo
		rows.Scan(&tag.id, &tag.tagName)
		tags[tag.tagName] = tag.id
	}
	return tags, nil
}

func (r *repository) ListDiariesByUserID(userID uint64, limit, offset uint64) ([]*model.Diary, error) {
	diaries := make([]*model.Diary, 0, limit)
	err := r.db.Select(&diaries,
		`SELECT id, name, user_id, updated_at FROM diary
			WHERE user_id = ?
			ORDER BY created_at DESC LIMIT ? OFFSET ?`,
		userID, limit, offset,
	)
	return diaries, err
}

func (r *repository) DeleteDiary(userID, diaryID uint64) (err error) {
	_, err = r.db.Exec(
		`DELETE FROM diary
			WHERE user_id = ? AND id = ?`,
		userID, diaryID,
	)
	if err != nil {
		return
	}
	_, err = r.db.Exec(
		`DELETE FROM article
			WHERE diary_id = ?`,
		diaryID,
	)
	return
}

func (r *repository) FindDiaryByID(diaryID uint64) (*model.Diary, error) {
	var diary model.Diary
	err := r.db.Get(
		&diary,
		`SELECT id, name, user_id FROM diary
			WHERE id = ? LIMIT 1`,
		diaryID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, diaryNotFoundError
		}
		return nil, err
	}
	return &diary, nil
}

func (r *repository) ListDiariesByIDs(diaryIDs []uint64) ([]*model.Diary, error) {
	if len(diaryIDs) == 0 {
		return nil, nil
	}
	diaries := make([]*model.Diary, 0, len(diaryIDs))
	query, args, err := sqlx.In(
		`SELECT id, name, user_id, updated_at FROM diary
			WHERE id IN (?)
			ORDER BY created_at DESC`, diaryIDs,
	)
	if err != nil {
		return nil, err
	}
	err = r.db.Select(&diaries, query, args...)
	return diaries, err
}

func (r *repository) ListDiariesByUserIDs(userIDs []uint64) (map[uint64][]*model.Diary, error) {
	if len(userIDs) == 0 {
		return nil, nil
	}
	query, args, err := sqlx.In(
		`SELECT id, name, user_id, updated_at FROM diary
		WHERE user_id IN (?)
		ORDER BY created_at DESC`,
		userIDs,
	)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	diaries := make(map[uint64][]*model.Diary)
	for rows.Next() {
		var diary model.Diary
		rows.Scan(&diary.ID, &diary.Name, &diary.UserID, &diary.UpdatedAt)
		diaries[diary.UserID] = append(diaries[diary.UserID], &diary)
	}
	return diaries, nil
}

func values(m map[string]uint64) []uint64 {
	vs := []uint64{}
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

func (r *repository) ListDiariesByTagIDs(tagIDs []uint64) (map[uint64][]*model.Diary, error) {
	if len(tagIDs) == 0 {
		return nil, nil
	}
	query, args, err := sqlx.In(
		`SELECT diary.id, diary.name, user_id, diary.updated_at, tag_id FROM diary
		JOIN diary_tag ON diary.id = diary_tag.diary_id
		WHERE tag_id IN (?)
		`, tagIDs,
	)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	diaries := make(map[uint64][]*model.Diary)
	for rows.Next() {
		var diary model.Diary
		var tagID uint64
		rows.Scan(&diary.ID, &diary.Name, &diary.UserID, &diary.UpdatedAt, &tagID)
		diaries[tagID] = append(diaries[tagID], &diary)
	}
	return diaries, nil
}

func (r *repository) getNoNullTags(diaryID uint64) ([]*model.Tag, error) {
	var rec []*model.TagRecord
	err := r.db.Select(&rec, `
		SELECT tag.id, tag_name, category_id FROM tag
			JOIN diary_tag ON tag.id = diary_tag.tag_id
			WHERE diary_id = ?
			`, diaryID,
	)
	if err != nil {
		return nil, err
	}
	var tagsOfDiary []*model.Tag
	for _, tag := range rec {
		if tag.CategoryID.Valid == true {
			tagsOfDiary = append(tagsOfDiary, &model.Tag{
				ID:         tag.ID,
				TagName:    tag.TagName,
				CategoryID: int(tag.CategoryID.Int64),
			})
		}
	}
	return tagsOfDiary, nil
}

func (r *repository) validateUserForDiaryMutation(diaryID, userID uint64) error {
	diary, err := r.FindDiaryByID(diaryID)
	if err != nil {
		return errors.New("diary not found")
	}
	if userID != diary.UserID {
		return errors.New("Unauthorized")
	}
	return nil
}

func (r *repository) ListCategories() []*model.Category {
	return instance.GetInstance()
}
