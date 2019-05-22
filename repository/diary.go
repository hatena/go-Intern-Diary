package repository

import (
	"database/sql"
	"time"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/jmoiron/sqlx"
)

var diaryNotFoundError = model.NotFoundError("diary")

func (r *repository) CreateNewDiary(userID uint64, name string, tags []string) (*model.Diary, error) { // Todo transaction and rollback
	now := time.Now()
	storedTagMap, err := r.getTagIDsByNames(tags)
	if err != nil {
		return nil, err
	}
	newTags := r.newTagFilter(tags, storedTagMap)
	newTagIds, err := r.insertNewTags(newTags, now)
	if err != nil {
		return nil, err
	}
	tagIds := append(values(storedTagMap), newTagIds...)

	newDiary, err := r.createNewDiary(userID, name, now)
	if err != nil {
		return nil, err
	}
	err = r.insertDiaryTags(newDiary.ID, tagIds, now)
	if err != nil {
		return nil, err
	}
	return newDiary, nil
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

func (r *repository) createNewDiary(userID uint64, name string, now time.Time) (*model.Diary, error) {
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
		WHERE id IN (?)
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

const DEFAULT_CATEGORY_ID = uint64(0)

func (r *repository) insertNewTags(newTagNames []string, now time.Time) ([]uint64, error) {
	// type Tag struct { // バルクインサートできない　https://github.com/jmoiron/sqlx/pull/285
	// 	id          uint64
	// 	tag_name    string
	// 	category_id uint64
	// 	updated_at  time.Time
	// 	created_at  time.Time
	// }
	// newTags := make([]*Tag, len(newTagNames))
	// for i, tagName := range newTagNames {
	// 	id, err := r.generateID()
	// 	if err != nil {
	// 		return []uint64{}, err
	// 	}
	// 	newTags[i] = &Tag{
	// 		id:          id,
	// 		tag_name:    tagName,
	// 		category_id: uint64(0), // Todo
	// 		updated_at:  now,
	// 		created_at:  now,
	// 	}
	// }
	// _, err := r.db.NamedExec(
	// 	`INSERT INTO tag
	// 		(id, tag_name, category_id, updated_at, created_at)
	// 		VALUES
	// 		(:id, :tag_name, :category_id, :updated_at, :created_at)
	// 	`, newTags,
	// )
	// ids := make([]uint64, len(newTags))
	// for i, newTag := range newTags {
	// 	ids[i] = newTag.id
	// }
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

func (r *repository) insertDiaryTags(diaryId uint64, tagIds []uint64, now time.Time) error {
	// type DiaryTag struct { // 同じくバルクインサートしたい
	// 	id         uint64
	// 	diary_id   uint64
	// 	tag_id     uint64
	// 	updated_at time.Time
	// 	created_at time.Time
	// }
	// newDiaryTags := make([]*DiaryTag, len(tagIds))
	// for i, tagId := range tagIds {
	// 	id, err := r.generateID()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	newDiaryTags[i] = &DiaryTag{
	// 		id:         id,
	// 		diary_id:   diaryId,
	// 		tag_id:     tagId,
	// 		updated_at: now,
	// 		created_at: now,
	// 	}
	// }
	// _, err := r.db.NamedExec(
	// 	`INSERT INTO diary_tag
	// 		(id, diary_id, tag_id, updated_at, created_at)
	// 		VALUES
	// 		(:id, :diary_id, :tag_id, :updated_at, :created_at)
	// 	`, newDiaryTags,
	// )
	for _, tagId := range tagIds {
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

func (r *repository) FindDiaryByID(diaryID, userID uint64) (*model.Diary, error) {
	var diary model.Diary
	err := r.db.Get(
		&diary,
		`SELECT id, name FROM diary
			WHERE id = ? AND user_id = ? LIMIT 1`,
		diaryID, userID,
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
		`SELECT id, name, user_id, updated_at, tag_id FROM diary
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
