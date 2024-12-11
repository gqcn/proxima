package user

import (
	"context"

	"proxima/app/gateway/utility"
	account "proxima/app/user/api/account/v1"

	"proxima/app/gateway/api/user/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	var (
		conn   = utility.UserClientConn()
		client = account.NewAccountClient(conn)
	)

	user, err := client.UserLogin(ctx, &account.UserLoginReq{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &v1.LoginRes{
		Token: user.GetToken(),
	}, nil
}
