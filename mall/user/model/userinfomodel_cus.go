package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func (m *defaultUserinfoModel) LoginFind(ctx context.Context, userid string, passwd string) (sql.Result, error) {
	userinfoUseridKey := fmt.Sprintf("%s%v", cacheUserinfoUseridPrefix, userid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("select 1 from %s where `userid` = ? and `password` = ?", m.table)
		return conn.ExecCtx(ctx, query, userid, passwd)
	}, userinfoUseridKey)
	return ret, err
}

func (m *defaultUserinfoModel) UpdatePasswd(ctx context.Context, userid string, passwd string) error {
	userinfoUseridKey := fmt.Sprintf("%s%v", cacheUserinfoUseridPrefix, userid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `password` = ? where `userid` = ?", m.table)
		return conn.ExecCtx(ctx, query, passwd, userid)
	}, userinfoUseridKey)
	return err
}
