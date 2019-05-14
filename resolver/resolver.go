package resolver

import (
	"context"
	"errors"
	"strconv"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

type Resolver interface {
	Visitor(context.Context) (*userResolver, error)
	GetUser(context.Context, struct{ UserID string }) (*userResolver, error)
	GetDiary(context.Context, struct{ DiaryID string }) (*diaryResolver, error)
	CreateDiary(context.Context, struct{ Name string }) (*diaryResolver, error)
	DeleteDiary(context.Context, struct{ DiaryID string }) (bool, error)
	PostArticle(context.Context, struct{ DiaryID, Title, Content string }) (*articleResolver, error)
	UpdateArticle(context.Context, struct{ ArticleID, Title, Content string }) (*articleResolver, error)
	DeleteArticle(context.Context, struct{ ArticleID, DiaryID string }) (bool, error)
}

func newResolver(app service.DiaryApp) Resolver {
	return &resolver{app: app}
}

type resolver struct {
	app service.DiaryApp
}

func currentUser(ctx context.Context) *model.User {
	return ctx.Value("user").(*model.User)
}

func (r *resolver) Visitor(ctx context.Context) (*userResolver, error) {
	if currentUser(ctx) == nil {
		return nil, errors.New("please login")
	}
	return &userResolver{currentUser(ctx)}, nil
}

func (r *resolver) GetUser(ctx context.Context, args struct{ UserID string }) (*userResolver, error) {
	userID, err := strconv.ParseUint(args.UserID, 10, 64)
	if err != nil {
		return nil, err
	}
	user, err := r.app.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return &userResolver{user}, nil
}

func (r *resolver) GetDiary(ctx context.Context, args struct{ DiaryID string }) (*diaryResolver, error) {
	diaryID, err := strconv.ParseUint(args.DiaryID, 10, 64)
	if err != nil {
		return nil, err
	}
	user := currentUser(ctx)
	diary, err := r.app.FindDiaryByID(diaryID, user.ID)
	if err != nil {
		return nil, err
	}
	if diary == nil {
		return nil, errors.New("diary not found")
	}
	return &diaryResolver{diary: diary}, nil
}

func (r *resolver) CreateDiary(ctx context.Context, args struct{ Name string }) (*diaryResolver, error) {
	user := currentUser(ctx)
	if user == nil {
		return nil, errors.New("user not found")
	}
	diary, err := r.app.CreateNewDiary(user.ID, args.Name)
	if err != nil {
		return nil, err
	}
	return &diaryResolver{diary: diary}, nil
}

func (r *resolver) DeleteDiary(ctx context.Context, args struct{ DiaryID string }) (bool, error) {
	user := currentUser(ctx)
	if user == nil {
		return false, errors.New("user not found")
	}
	diaryID, err := strconv.ParseUint(args.DiaryID, 10, 64)
	if err != nil {
		return false, err
	}
	err = r.app.DeleteDiary(user.ID, diaryID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *resolver) PostArticle(ctx context.Context, args struct{ DiaryID, Title, Content string }) (*articleResolver, error) {
	user := currentUser(ctx)
	if user == nil {
		return nil, errors.New("user not found")
	}
	diaryID, err := strconv.ParseUint(args.DiaryID, 10, 64)
	if err != nil {
		return nil, err
	}
	article, err := r.app.CreateNewArticle(diaryID, args.Title, args.Content)
	if err != nil {
		return nil, err
	}
	return &articleResolver{article: article}, nil
}

func (r *resolver) DeleteArticle(ctx context.Context, args struct{ ArticleID, DiaryID string }) (bool, error) {
	user := currentUser(ctx)
	if user == nil {
		return false, errors.New("user not found")
	}
	articleID, err := strconv.ParseUint(args.ArticleID, 10, 64)
	if err != nil {
		return false, err
	}
	err = r.app.DeleteArticle(articleID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *resolver) UpdateArticle(ctx context.Context, args struct{ ArticleID, Title, Content string }) (*articleResolver, error) {
	user := currentUser(ctx)
	if user == nil {
		return nil, errors.New("user not found")
	}
	articleID, err := strconv.ParseUint(args.ArticleID, 10, 64)
	if err != nil {
		return nil, err
	}
	article, err := r.app.UpdateArticle(articleID, args.Title, args.Content)
	if err != nil {
		return nil, err
	}
	return &articleResolver{article: article}, nil
}
