package logic

import (
	"context"
	"encoding/json"

	"go_zero/study_api/user/api_jwt/internal/svc"
	"go_zero/study_api/user/api_jwt/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	//獲取token值
	userid := l.ctx.Value("userId").(json.Number)

	//遇到不知道的類型錯誤檢查方式
	//fmt.Printf("%v %T",userid,userid)
	uuid, _ := userid.Int64()

	username := l.ctx.Value("username").(string)
	return &types.UserInfoResponse{
		UserId:   uint(uuid),
		Username: username,
	}, nil
}
