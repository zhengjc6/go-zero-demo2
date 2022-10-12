package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"strings"

	"go-zero-demo2/mall/user/api/internal/svc"
	"go-zero-demo2/mall/user/api/internal/types"
	"go-zero-demo2/mall/user/model"

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

func (l *LoginLogic) Login(req *types.UserLoginRequest) (resp *types.UserLoginReply, err error) {
	if len(strings.TrimSpace(req.Userid)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("invliad param")
	}

	m := md5.New()
	io.WriteString(m, req.Password)
	cryptstr := hex.EncodeToString(m.Sum(nil))
	_, sqlerr := l.svcCtx.UserinfoModel.LoginFind(l.ctx, req.Userid, cryptstr)
	switch sqlerr {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("userid not exist")
	default:
		return nil, sqlerr
	}
	//create jwt
	jwttoekn, err := getJwtToken(l.svcCtx.Config, req.Userid)
	if err != nil {
		return nil, errors.New("create jwt fail")
	}
	//log
	return &types.UserLoginReply{AccessToken: jwttoekn, AccessExpire: l.svcCtx.Config.Auth.AccessExpire}, nil

}
