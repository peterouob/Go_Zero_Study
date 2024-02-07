package logic

import (
	"context"
	"go_zero/common/jwt"

	"go_zero/study_api/user/api_jwt/internal/svc"
	"go_zero/study_api/user/api_jwt/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	//獲取配置文件
	auth := l.svcCtx.Config.Auth
	token, err := jwt.GetToken(jwt.JwtPayload{
		UserId:   1,
		Username: "peter",
		Role:     1,
	}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return "", err
	}
	return token, nil
}
