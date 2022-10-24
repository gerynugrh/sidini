package forum

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
	Name string `json:"name"`
}

func (u UseCase) GetList(ctx context.Context, param GetListParam) ([]Forum, error) {
	forums := make([]Forum, 0)

	err := common.Dao(ctx).DB().
		NewQuery("select * from forums where name like %{:name}%").
		Bind(dbx.Params{"name": param.Name}).
		All(&forums)
	if err != nil {
		return nil, fmt.Errorf("db.NewQuery: %w", err)
	}

	return forums, nil
}
