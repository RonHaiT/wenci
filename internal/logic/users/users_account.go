package users

import (
	"context"
	"strings"
	"time"
	"wenci/internal/consts"
	"wenci/internal/dao"
	"wenci/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	Id       uint
	Username string

	jwt.RegisteredClaims
}

func (u *Users) Login(ctx context.Context, username string, password string) (tokenString string, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("username", username).Scan(&user)
	if err != nil {
		return "", gerror.New("用户名或密码错误")
	}
	if user.Id == 0 {
		return "", gerror.New("用户不存在")
	}
	if user.Password != u.encryptPassword(password) {
		return "", gerror.New("用户名或密码错误")
	}
	// 生成token
	uc := &jwtClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString([]byte(consts.JwtKey))
}

func (u *Users) Info(ctx context.Context) (user *entity.Users, err error) {
	raw := g.RequestFromCtx(ctx).Header.Get("Authorization")
	tokenString := strings.TrimSpace(raw)
	if tokenString == "" {
		return nil, gerror.New("未登录或登录已过期")
	}

	const bearerPrefix = "Bearer "
	if strings.HasPrefix(tokenString, bearerPrefix) {
		tokenString = strings.TrimSpace(tokenString[len(bearerPrefix):])
	}

	tokenClaims, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})
	if err != nil || !tokenClaims.Valid {
		return nil, gerror.New("未登录或登录已过期")
	}

	claims, ok := tokenClaims.Claims.(*jwtClaims)
	if !ok {
		return nil, gerror.New("未登录或登录已过期")
	}

	err = dao.Users.Ctx(ctx).Where("id", claims.Id).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Id == 0 {
		return nil, gerror.New("用户不存在")
	}
	return user, nil
}

func (u *Users) GetUid(ctx context.Context) (uint, error) {
	user, err := u.Info(ctx)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
