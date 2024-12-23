package account

import (
    "context"
    "time"

    "proxima/app/user/internal/dao"
    "proxima/app/user/internal/model/do"
    "proxima/app/user/internal/model/entity"

    "github.com/gogf/gf/v2/errors/gerror"
    "github.com/golang-jwt/jwt/v5"
)

var (
    // In production, this should be loaded from configuration
    jwtSecret       = []byte("jwt_secret_key")
    tokenExpiration = 24 * time.Hour
)

type Account struct{}

type jwtClaims struct {
    UserID int64 `json:"user_id"`
    jwt.RegisteredClaims
}

func New() *Account {
    return &Account{}
}

type RegisterInput struct {
    Passport string
    Password string
    Email    string
}

func (*Account) Register(ctx context.Context, in RegisterInput) (id int64, err error) {
    return dao.User.Ctx(ctx).InsertAndGetId(do.User{
        Passport: in.Passport,
        Password: in.Password,
        Email:    in.Email,
    })
}

func (*Account) Login(ctx context.Context, passport, password string) (token string, err error) {
    var user *entity.User
    err = dao.User.Ctx(ctx).Where(do.User{
        Passport: passport,
        Password: password,
    }).Scan(&user)
    if err != nil {
        return "", err
    }
    if user == nil {
        return "", gerror.New("invalid credentials")
    }

    var (
        now    = time.Now()
        claims = jwtClaims{
            UserID: int64(user.Id),
            RegisteredClaims: jwt.RegisteredClaims{
                ExpiresAt: jwt.NewNumericDate(now.Add(tokenExpiration)),
                IssuedAt:  jwt.NewNumericDate(now),
                NotBefore: jwt.NewNumericDate(now),
            },
        }
    )
    token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
    if err != nil {
        return "", gerror.Wrapf(err, `jwt claims failed`)
    }
    return token, nil
}

func (*Account) Info(ctx context.Context, token string) (user *entity.User, err error) {
    claims := &jwtClaims{}
    parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        return nil, gerror.Wrapf(err, `jwt token parsing failed`)
    }
    if !parsedToken.Valid {
        return nil, gerror.New("invalid token")
    }

    err = dao.User.Ctx(ctx).Where("id", claims.UserID).Scan(&user)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, gerror.New("user not found")
    }
    return
}
