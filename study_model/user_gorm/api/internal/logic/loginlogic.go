package logic

import (
	"context"
	"fmt"
	"go_zero/study_model/user_gorm/api/internal/svc"
	"go_zero/study_model/user_gorm/api/internal/types"
	"go_zero/study_model/user_gorm/models"

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
	//err = l.svcCtx.DB.Create(&models.GUserModel{
	//	Username: "defer",
	//	Password: "123456",
	//}).Error
	user := &models.UserModel{}
	err = l.svcCtx.DB.Find(user).Error
	if err != nil {
		return "", err
	}
	fmt.Println(user)
	return "success", nil
}
