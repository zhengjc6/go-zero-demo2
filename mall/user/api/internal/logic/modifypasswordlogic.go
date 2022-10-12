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

type ModifypasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifypasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifypasswordLogic {
	return &ModifypasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifypasswordLogic) Modifypassword(req *types.UserPasswordModifyRequest) (resp *types.UserPasswordModifyReply, err error) {
	if len(strings.TrimSpace(req.Userid)) == 0 ||
		len(strings.TrimSpace(req.OldPassword)) == 0 ||
		len(strings.TrimSpace(req.NewPassword)) == 0 {
		return nil, errors.New("invliad param")
	}

	m := md5.New()
	io.WriteString(m, req.OldPassword)
	cryptstr := hex.EncodeToString(m.Sum(nil))
	_, sqlerr := l.svcCtx.UserinfoModel.LoginFind(l.ctx, req.Userid, cryptstr)
	switch sqlerr {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("userid not exist")
	default:
		return nil, sqlerr
	}
	//update new password

	m.Reset()
	io.WriteString(m, req.NewPassword)
	cryptstr = hex.EncodeToString(m.Sum(nil))
	sqlerr = l.svcCtx.UserinfoModel.UpdatePasswd(l.ctx, req.Userid, req.NewPassword)
	if sqlerr != nil {
		return nil, errors.New("Modifypassword fail internal error")
	}
	//create new jwt
	_, err = getJwtToken(l.svcCtx.Config, req.Userid)
	if err != nil {
		return nil, errors.New("create jwt fail")
	}
	//log
	return &types.UserPasswordModifyReply{}, nil
}
