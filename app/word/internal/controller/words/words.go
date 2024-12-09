package words

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
	"proxima/app/word/api/pbentity"
	v1 "proxima/app/word/api/words/v1"
	"proxima/app/word/internal/logic/words"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedWordsServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterWordsServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := words.Create(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: uint32(id)}, nil
}

func (*Controller) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	data, err := words.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetRes{
		Words: &pbentity.Words{
			Id:                 uint32(data.Id),
			Uid:                uint32(data.Uid),
			Word:               data.Word,
			Definition:         data.Definition,
			ExampleSentence:    data.ExampleSentence,
			ChineseTranslation: data.ChineseTranslation,
			Pronunciation:      data.Pronunciation,
			CreatedAt:          timestamppb.New(data.CreatedAt.Time),
			UpdatedAt:          timestamppb.New(data.CreatedAt.Time),
		},
	}, nil
}
