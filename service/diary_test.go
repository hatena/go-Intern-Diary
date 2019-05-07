package service

import (
	"testing"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/stretchr/testify/assert"
)

func createTestDiary(app DiaryApp, userID uint64) *model.Diary {
	name := "test name " + randomString()
	diary, err := app.CreateNewDiary(userID, name)
	if err != nil {
		panic(err)
	}
	return diary
}

func TestDiaryApp_ListDiariesByUserID(t *testing.T) {
	app := newApp()
	defer closeApp(app)

	name := "test name " + randomString()
	password := randomString() + randomString()
	_ = app.CreateNewUser(name, password)
	user, _ := app.FindUserByName(name)
	assert.Equal(t, name, user.Name)

	diaries := make([]*model.Diary, 5, 5)
	for i := range diaries {
		diaries[i] = createTestDiary(app, user.ID)
	}

	diariesFound, err := app.ListDiariesByUserID(user.ID, 1, 10)
	assert.NoError(t, err)
	assert.Len(t, diariesFound, 5)
}
