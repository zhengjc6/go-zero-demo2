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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	if len(strings.TrimSpace(req.Userid)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("invliad param")
	}

	_, sqlerr := l.svcCtx.UserinfoModel.FindOne(l.ctx, req.Userid)
	switch sqlerr {
	case nil:
		return nil, errors.New("userid exist")
	case model.ErrNotFound:
		//pass
	default:
		return nil, sqlerr
	}
	m := md5.New()
	io.WriteString(m, req.Password)
	cryptstr := hex.EncodeToString(m.Sum(nil))
	_, sqlerr = l.svcCtx.UserinfoModel.Insert(l.ctx, &model.Userinfo{Userid: req.Userid, Password: cryptstr})
	if sqlerr != nil {
		return nil, errors.New("register fail internal error")
	}
	//log
	return &types.UserRegisterReply{}, nil
}
