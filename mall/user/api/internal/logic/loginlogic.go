package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"strings"

	errorx "go-zero-demo2/mall/user/api/internal/common/userapierror"
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
	ret, sqlerr := l.svcCtx.UserinfoModel.FindOne(l.ctx, req.Userid)
	if sqlerr != nil {
		if sqlerr == model.ErrNotFound {
			return nil, errorx.NewDefaultError("user not exit1")
		}
		return nil, errorx.NewDefaultError("internal db error")
	}
	if ret.Password != cryptstr {
		return nil, errorx.NewDefaultError("user not exit2")
	}
	//create jwt
	jwttoekn, err := getJwtToken(l.svcCtx.Config, req.Userid)
	if err != nil {
		return nil, errorx.NewDefaultError("create jwt fail")
	}
	//log
	return &types.UserLoginReply{AccessToken: jwttoekn, AccessExpire: l.svcCtx.Config.Auth.AccessExpire}, nil

}
