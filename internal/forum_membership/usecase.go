package forum_membership

import (
	"context"
	"fmt"
	"github.com/gerynugrh/ubiquitous-octo-waffle/internal/common"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/models"
)

type UseCase struct{}

func NewUseCase() UseCase {
	return UseCase{}
}

type GetListParam struct {
	ForumID uint64
}

func (u UseCase) GetList(ctx context.Context, param GetListParam) ([]*models.User, error) {
	userIDs := make([]string, 0)

	err := common.Dao(ctx).DB().
		NewQuery("select user_id from forum_memberships where forum_id = {:forumID}").
		Bind(dbx.Params{"forumID": param.ForumID}).
		All(&userIDs)
	if err != nil {
		return nil, fmt.Errorf("db.NewQuery: %w", err)
	}

	users := make([]*models.User, 0)
	err = common.Dao(ctx).DB().
		NewQuery("select * from _users where id in ({:userIDs})").
		Bind(dbx.Params{"userIDs": userIDs}).
		All(&users)
	if err != nil {
		return nil, fmt.Errorf("db.NewQuery: %w", err)
	}

	return users, nil
}
