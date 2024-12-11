package words

import (
	"context"

	"proxima/app/gateway/utility"
	words "proxima/app/word/api/words/v1"

	"proxima/app/gateway/api/words/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	var (
		conn   = utility.WordClientConn()
		client = words.NewWordsClient(conn)
	)

	ctx, cancel := context.WithTimeout(ctx, utility.Timeout)
	defer cancel()

	_, err = client.Create(ctx, &words.CreateReq{
		Uid:        1,
		Word:       req.Word,
		Definition: req.Definition,
	})

	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{}, nil
}
