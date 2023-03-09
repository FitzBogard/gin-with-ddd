package infrastructure

import (
	"context"
	"go-ddd/domain"
	"go-ddd/util"
)

type (
	ArticleRepoImpl struct {
	}
	Article struct {
		ID          int64
		Title       string
		UserID      int64
		PublishAt   int64
		LastReplyAt int64
	}
	ArticleRepo interface {
		FindByID(ctx context.Context, id int64) Article
	}
)

func (a *ArticleRepoImpl) FindByID(ctx context.Context, id int64) Article {
	var article *domain.Article
	util.Orm.First(&article, id)
	return Article{
		ID:          article.ID,
		Title:       article.Title,
		UserID:      article.UserID,
		PublishAt:   article.PublishAt,
		LastReplyAt: article.LastReplyAt,
	}
}
