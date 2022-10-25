package message

import (
	"context"
	"fmt"
	"github.com/gerynugrh/ubiquitous-octo-waffle/internal/common"
	"github.com/pocketbase/dbx"
)

type UseCase struct{}

func NewUseCase() UseCase {
	return UseCase{}
}

type GetListParam struct {
	ForumID uint64 `json:"forumID"`
}

func (u UseCase) GetList(ctx context.Context, param GetListParam) ([]Message, error) {
	messages := make([]Message, 0)
	err := common.Dao(ctx).DB().
		NewQuery("select * from messages where forum_id = ({:forumID})").
		Bind(dbx.Params{"forumID": param.ForumID}).
		All(&messages)
	if err != nil {
		return nil, fmt.Errorf("db.NewQuery: %w", err)
	}

	return messages, nil
}
