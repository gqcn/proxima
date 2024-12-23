package account

import (
    "context"

    v1 "proxima/app/user/api/account/v1"
    "proxima/app/user/internal/logic/account"

    "github.com/gogf/gf/contrib/rpc/grpcx/v2"
    "github.com/gogf/gf/v2/util/gconv"
)

type Controller struct {
    v1.UnimplementedAccountServer
    account *account.Account
}

func Register(s *grpcx.GrpcServer) {
    v1.RegisterAccountServer(s.Server, &Controller{
        account: account.New(),
    })
}

func (c *Controller) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
    id, err := c.account.Register(ctx, account.RegisterInput{
        Passport: req.Passport,
        Password: req.Password,
        Email:    req.Email,
    })
    if err != nil {
        return nil, err
    }
    return &v1.RegisterRes{
        Id: int32(id),
    }, nil
}

func (c *Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
    token, err := c.account.Login(ctx, req.Passport, req.Passport)
    if err != nil {
        return nil, err
    }
    return &v1.LoginRes{
        Token: token,
    }, nil
}

func (c *Controller) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
    entityData, err := c.account.Info(ctx, req.Token)
    if err != nil {
        return nil, err
    }
    res = &v1.InfoRes{}
    err = gconv.Scan(entityData, &res.User)
    return
}
