// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	ordersFieldNames          = builder.RawFieldNames(&Orders{})
	ordersRows                = strings.Join(ordersFieldNames, ",")
	ordersRowsExpectAutoSet   = strings.Join(stringx.Remove(ordersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	ordersRowsWithPlaceHolder = strings.Join(stringx.Remove(ordersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheOrdersIdPrefix = "cache:orders:id:"
)

type (
	ordersModel interface {
		Insert(ctx context.Context, data *Orders) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Orders, error)
		Update(ctx context.Context, data *Orders) error
		Delete(ctx context.Context, id int64) error
	}

	defaultOrdersModel struct {
		sqlc.CachedConn
		table string
	}

	Orders struct {
		Id         int64           `db:"id"`
		UserID     int64           `db:"userID"`
		CreateTime int64           `db:"createTime"`
		Name       string          `db:"name"`
		Total      sql.NullFloat64 `db:"total"`
	}
)

func newOrdersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultOrdersModel {
	return &defaultOrdersModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`orders`",
	}
}

func (m *defaultOrdersModel) Delete(ctx context.Context, id int64) error {
	ordersIdKey := fmt.Sprintf("%s%v", cacheOrdersIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, ordersIdKey)
	return err
}

func (m *defaultOrdersModel) FindOne(ctx context.Context, id int64) (*Orders, error) {
	ordersIdKey := fmt.Sprintf("%s%v", cacheOrdersIdPrefix, id)
	var resp Orders
	err := m.QueryRowCtx(ctx, &resp, ordersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ordersRows, m.table)
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

func (m *defaultOrdersModel) Insert(ctx context.Context, data *Orders) (sql.Result, error) {
	ordersIdKey := fmt.Sprintf("%s%v", cacheOrdersIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, ordersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserID, data.CreateTime, data.Name, data.Total)
	}, ordersIdKey)
	return ret, err
}

func (m *defaultOrdersModel) Update(ctx context.Context, data *Orders) error {
	ordersIdKey := fmt.Sprintf("%s%v", cacheOrdersIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ordersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserID, data.CreateTime, data.Name, data.Total, data.Id)
	}, ordersIdKey)
	return err
}

func (m *defaultOrdersModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheOrdersIdPrefix, primary)
}

func (m *defaultOrdersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ordersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultOrdersModel) tableName() string {
	return m.table
}
