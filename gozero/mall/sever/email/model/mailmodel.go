package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"strings"

	sqlBuilder "github.com/didi/gendry/builder"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	mailFieldNames          = builder.RawFieldNames(&Mail{})
	mailRows                = strings.Join(mailFieldNames, ",")
	mailRowsExpectAutoSet   = strings.Join(stringx.Remove(mailFieldNames, "`id`"), ",")
	mailRowsWithPlaceHolder = strings.Join(stringx.Remove(mailFieldNames, "`id`"), "=?,") + "=?"

	cacheMailIdPrefix = "cache:mail:id:"
)

type (
	MailModel interface {
		Insert(ctx context.Context, data *Mail) (sql.Result, error)
		InsertNoCache(ctx context.Context, data *Mail) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Mail, error)
		FindsNoCache(ctx context.Context, pageNum, pageSize int, where map[string]interface{}) ([]*Mail, error)
		Update(ctx context.Context, newData *Mail) error
		Delete(ctx context.Context, id int64) error
	}

	defaultMailModel struct {
		sqlc.CachedConn
		table string
	}

	Mail struct {
		Id        int64          `db:"id"`
		Name      sql.NullString `db:"name"`       // 收件人姓名
		ToEmail   sql.NullString `db:"to_email"`   // 收件人邮箱
		IsSend    int            `db:"is_send"`    // 1=已经发送，0=没有发送
		CreatedAt sql.NullInt64  `db:"created_at"` // 创建时间
		UpdateAt  sql.NullInt64  `db:"update_at"`  // 更新时间
	}
)

// NewMailModel 改造一下，返回接口
func NewMailModel(conn sqlx.SqlConn, c cache.CacheConf) MailModel {
	return &defaultMailModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`mail`",
	}

}

// InsertNoCache 插入不删除缓存
func (m *defaultMailModel) InsertNoCache(ctx context.Context, data *Mail) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s(%s) values (?,?,?,?,?)", m.table, mailRowsExpectAutoSet)
	ret, err := m.ExecNoCacheCtx(ctx, query, data.Name, data.ToEmail, data.IsSend, data.CreatedAt, data.UpdateAt)
	return ret, err
}

// Insert 插入需要删除缓存数据
func (m *defaultMailModel) Insert(ctx context.Context, data *Mail) (sql.Result, error) {
	mailIdKey := fmt.Sprintf("%s%v", cacheMailIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, mailRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.ToEmail, data.IsSend, data.CreatedAt, data.UpdateAt)
	}, mailIdKey)
	return ret, err

}

// FindOne 查找加入缓存
func (m *defaultMailModel) FindOne(ctx context.Context, id int64) (*Mail, error) {
	mailIdKey := fmt.Sprintf("%s%v", cacheMailIdPrefix, id)
	var resp Mail
	err := m.QueryRowCtx(ctx, &resp, mailIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", mailRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindsNoCache 查询多条(分页）-不带缓存
func (m *defaultMailModel) FindsNoCache(ctx context.Context, pageNum, pageSize int, where map[string]interface{}) ([]*Mail, error) {
	// 分页查询
	where["_limit"] = []uint{uint(pageSize * (pageNum - 1)), uint(pageSize)}
	// sql 构造器
	query, values, err := sqlBuilder.BuildSelect(m.table, where, mailFieldNames)
	if err != nil {
		return nil, err
	}
	var resp []*Mail
	//
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Update 更新，需要更新缓存
func (m *defaultMailModel) Update(ctx context.Context, data *Mail) error {
	mailIdKey := fmt.Sprintf("%s%v", cacheMailIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, mailRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.ToEmail, data.IsSend, data.CreatedAt, data.UpdateAt, data.Id)
	}, mailIdKey)
	return err
}

func (m *defaultMailModel) Delete(ctx context.Context, id int64) error {
	mailIdKey := fmt.Sprintf("%s%v", cacheMailIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, mailIdKey)
	return err
}

func (m *defaultMailModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMailIdPrefix, primary)
}

func (m *defaultMailModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", mailRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultMailModel) tableName() string {
	return m.table
}
