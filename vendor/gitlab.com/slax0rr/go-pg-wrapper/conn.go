/*
Copyright © 2020 Tomaz Lovrec <tomaz.lovrec@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package pgwrapper

import (
	"context"
	"io"
	"time"

	"github.com/go-pg/pg/v9"
	pgorm "github.com/go-pg/pg/v9/orm"
	"gitlab.com/slax0rr/go-pg-wrapper/orm"
)

type Conn interface {
	AddQueryHook(hook pg.QueryHook)
	Begin() (Tx, error)
	Close() error
	Context() context.Context
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error)
	CreateComposite(model interface{}, opt *pgorm.CreateCompositeOptions) error
	CreateTable(model interface{}, opt *pgorm.CreateTableOptions) error
	Delete(model interface{}) error
	DropComposite(model interface{}, opt *pgorm.DropCompositeOptions) error
	DropTable(model interface{}, opt *pgorm.DropTableOptions) error
	Exec(query interface{}, params ...interface{}) (res pg.Result, err error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	ExecOne(query interface{}, params ...interface{}) (pg.Result, error)
	ExecOneContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	ForceDelete(model interface{}) error
	Formatter() pgorm.QueryFormatter
	Insert(model ...interface{}) error
	Model(model ...interface{}) orm.Query
	ModelContext(c context.Context, model ...interface{}) orm.Query
	Param(param string) interface{}
	PoolStats() *pg.PoolStats
	Prepare(q string) (Stmt, error)
	Query(model, query interface{}, params ...interface{}) (res pg.Result, err error)
	QueryContext(c context.Context, model, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (pg.Result, error)
	RunInTransaction(fn func(Tx) error) error
	Select(model interface{}) error
	Update(model interface{}) error
	WithContext(ctx context.Context) Conn
	WithParam(param string, value interface{}) Conn
	WithTimeout(d time.Duration) Conn
}

type ConnWrap struct {
	conn *pg.Conn
}

func NewConn(conn *pg.Conn) *ConnWrap {
	return &ConnWrap{conn}
}

func (c *ConnWrap) AddQueryHook(hook pg.QueryHook) {
	c.conn.AddQueryHook(hook)
}

func (c *ConnWrap) Begin() (Tx, error) {
	tx, err := c.conn.Begin()
	if err != nil {
		return nil, err
	}
	return NewTx(tx), nil
}

func (c *ConnWrap) Close() error {
	return c.conn.Close()
}

func (c *ConnWrap) Context() context.Context {
	return c.conn.Context()
}

func (c *ConnWrap) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return c.conn.CopyFrom(r, query, params...)
}

func (c *ConnWrap) CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return c.conn.CopyTo(w, query, params...)
}

func (c *ConnWrap) CreateComposite(model interface{}, opt *pgorm.CreateCompositeOptions) error {
	return c.conn.CreateComposite(model, opt)
}

func (c *ConnWrap) CreateTable(model interface{}, opt *pgorm.CreateTableOptions) error {
	return c.conn.CreateTable(model, opt)
}

func (c *ConnWrap) Delete(model interface{}) error {
	return c.conn.Delete(model)
}

func (c *ConnWrap) DropComposite(model interface{}, opt *pgorm.DropCompositeOptions) error {
	return c.conn.DropComposite(model, opt)
}

func (c *ConnWrap) DropTable(model interface{}, opt *pgorm.DropTableOptions) error {
	return c.conn.DropTable(model, opt)
}

func (c *ConnWrap) Exec(query interface{}, params ...interface{}) (res pg.Result, err error) {
	return c.conn.Exec(query, params...)
}

func (c *ConnWrap) ExecContext(ctx context.Context, query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.ExecContext(ctx, query, params...)
}

func (c *ConnWrap) ExecOne(query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.ExecOne(query, params...)
}

func (c *ConnWrap) ExecOneContext(ctx context.Context, query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.ExecOneContext(ctx, query, params...)
}

func (c *ConnWrap) ForceDelete(model interface{}) error {
	return c.conn.ForceDelete(model)
}

func (c *ConnWrap) Formatter() pgorm.QueryFormatter {
	return c.conn.Formatter()
}

func (c *ConnWrap) Insert(model ...interface{}) error {
	return c.conn.Insert(model...)
}

func (c *ConnWrap) Model(model ...interface{}) orm.Query {
	return orm.NewQuery(c.conn.Model(model...))
}

func (c *ConnWrap) ModelContext(ctx context.Context, model ...interface{}) orm.Query {
	return orm.NewQuery(c.conn.ModelContext(ctx, model...))
}

func (c *ConnWrap) Param(param string) interface{} {
	return c.conn.Param(param)
}

func (c *ConnWrap) PoolStats() *pg.PoolStats {
	return c.conn.PoolStats()
}

func (c *ConnWrap) Prepare(q string) (Stmt, error) {
	stmt, err := c.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	return NewStmt(stmt), nil
}

func (c *ConnWrap) Query(model, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return c.conn.Query(model, query, params...)
}

func (c *ConnWrap) QueryContext(ctx context.Context, model, query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.QueryContext(ctx, model, query, params...)
}

func (c *ConnWrap) QueryOne(model, query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.QueryOne(model, query, params...)
}

func (c *ConnWrap) QueryOneContext(ctx context.Context, model, query interface{}, params ...interface{}) (pg.Result, error) {
	return c.conn.QueryOneContext(ctx, model, query, params...)
}

func (c *ConnWrap) RunInTransaction(fn func(Tx) error) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return err
	}
	return NewTx(tx).RunInTransaction(fn)
}

func (c *ConnWrap) Select(model interface{}) error {
	return c.conn.Select(model)
}

func (c *ConnWrap) Update(model interface{}) error {
	return c.conn.Update(model)
}

func (c *ConnWrap) WithContext(ctx context.Context) Conn {
	conn := c.conn.WithContext(ctx)
	return NewConn(conn)
}

func (c *ConnWrap) WithParam(param string, value interface{}) Conn {
	conn := c.conn.WithParam(param, value)
	return NewConn(conn)
}

func (c *ConnWrap) WithTimeout(d time.Duration) Conn {
	conn := c.conn.WithTimeout(d)
	return NewConn(conn)
}
