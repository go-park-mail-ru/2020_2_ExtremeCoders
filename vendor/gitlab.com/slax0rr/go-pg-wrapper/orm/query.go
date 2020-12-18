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
package orm

import (
	"context"
	"io"

	"github.com/go-pg/pg/v9"
	pgorm "github.com/go-pg/pg/v9/orm"
)

type Query interface {
	AllWithDeleted() Query
	AppendQuery(fmter pgorm.QueryFormatter, b []byte) ([]byte, error)
	Apply(fn func(Query) (Query, error)) Query
	Clone() Query
	Column(columns ...string) Query
	ColumnExpr(expr string, params ...interface{}) Query
	Context(c context.Context) Query
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (pgorm.Result, error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (pgorm.Result, error)
	Count() (int, error)
	CountEstimate(threshold int) (int, error)
	CreateTable(opt *pgorm.CreateTableOptions) error
	DB(db *pg.DB) Query
	Delete(values ...interface{}) (pgorm.Result, error)
	Deleted() Query
	Distinct() Query
	DistinctOn(expr string, params ...interface{}) Query
	DropTable(opt *pgorm.DropTableOptions) error
	Except(other Query) Query
	ExceptAll(other Query) Query
	ExcludeColumn(columns ...string) Query
	Exec(query interface{}, params ...interface{}) (pgorm.Result, error)
	ExecOne(query interface{}, params ...interface{}) (pgorm.Result, error)
	Exists() (bool, error)
	First() error
	For(s string, params ...interface{}) Query
	ForEach(fn interface{}) error
	ForceDelete(values ...interface{}) (pgorm.Result, error)
	Group(columns ...string) Query
	GroupExpr(group string, params ...interface{}) Query
	Having(having string, params ...interface{}) Query
	Insert(values ...interface{}) (pgorm.Result, error)
	Intersect(other Query) Query
	IntersectAll(other Query) Query
	Join(join string, params ...interface{}) Query
	JoinOn(condition string, params ...interface{}) Query
	JoinOnOr(condition string, params ...interface{}) Query
	Last() error
	Limit(n int) Query
	Model(model ...interface{}) Query
	New() Query
	Offset(n int) Query
	OnConflict(s string, params ...interface{}) Query
	Order(orders ...string) Query
	OrderExpr(order string, params ...interface{}) Query
	Query(model, query interface{}, params ...interface{}) (pgorm.Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (pgorm.Result, error)
	Relation(name string, apply ...func(Query) (Query, error)) Query
	Returning(s string, params ...interface{}) Query
	Select(values ...interface{}) error
	SelectAndCount(values ...interface{}) (count int, firstErr error)
	SelectAndCountEstimate(threshold int, values ...interface{}) (count int, firstErr error)
	SelectOrInsert(values ...interface{}) (inserted bool, _ error)
	Set(set string, params ...interface{}) Query
	Table(tables ...string) Query
	TableExpr(expr string, params ...interface{}) Query
	TableModel() pgorm.TableModel
	Union(other Query) Query
	UnionAll(other Query) Query
	Update(scan ...interface{}) (pgorm.Result, error)
	UpdateNotZero(scan ...interface{}) (pgorm.Result, error)
	Value(column string, value string, params ...interface{}) Query
	Where(condition string, params ...interface{}) Query
	WhereGroup(fn func(Query) (Query, error)) Query
	WhereIn(where string, slice interface{}) Query
	WhereInMulti(where string, values ...interface{}) Query
	WhereNotGroup(fn func(Query) (Query, error)) Query
	WhereOr(condition string, params ...interface{}) Query
	WhereOrGroup(fn func(Query) (Query, error)) Query
	WhereOrNotGroup(fn func(Query) (Query, error)) Query
	WherePK() Query
	WhereStruct(strct interface{}) Query
	With(name string, subq Query) Query
	WithDelete(name string, subq Query) Query
	WithInsert(name string, subq Query) Query
	WithUpdate(name string, subq Query) Query
	WrapWith(name string) Query
}

type QueryWrap struct {
	query *pgorm.Query
}

func NewQuery(query *pgorm.Query) *QueryWrap {
	return &QueryWrap{query}
}

func (q *QueryWrap) AllWithDeleted() Query {
	return NewQuery(q.query.AllWithDeleted())
}

func (q *QueryWrap) AppendQuery(fmter pgorm.QueryFormatter, b []byte) ([]byte, error) {
	return q.query.AppendQuery(fmter, b)
}

func (q *QueryWrap) Apply(fn func(Query) (Query, error)) Query {
	qq, err := fn(q)
	if err != nil {
		// Call original apply to with a function returning the above error to set `stickyErr` in go-pg/orm
		q.query.Apply(func(pgquery *pgorm.Query) (*pgorm.Query, error) {
			return nil, err
		})
		return q
	}
	return qq
}

func (q *QueryWrap) Clone() Query {
	return NewQuery(q.query.Clone())
}

func (q *QueryWrap) Column(columns ...string) Query {
	return NewQuery(q.query.Column(columns...))
}

func (q *QueryWrap) ColumnExpr(expr string, params ...interface{}) Query {
	return NewQuery(q.query.ColumnExpr(expr, params...))
}

func (q *QueryWrap) Context(c context.Context) Query {
	return NewQuery(q.query.Context(c))
}

func (q *QueryWrap) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (pgorm.Result, error) {
	return q.query.CopyFrom(r, query, params)
}

func (q *QueryWrap) CopyTo(w io.Writer, query interface{}, params ...interface{}) (pgorm.Result, error) {
	return q.query.CopyTo(w, query, params...)
}

func (q *QueryWrap) Count() (int, error) {
	return q.query.Count()
}

func (q *QueryWrap) CountEstimate(threshold int) (int, error) {
	return q.query.CountEstimate(threshold)
}

func (q *QueryWrap) CreateTable(opt *pgorm.CreateTableOptions) error {
	return q.query.CreateTable(opt)
}

func (q *QueryWrap) DB(db *pg.DB) Query {
	return NewQuery(q.query.DB(db))
}

func (q *QueryWrap) Delete(values ...interface{}) (pgorm.Result, error) {
	return q.query.Delete(values...)
}

func (q *QueryWrap) Deleted() Query {
	return NewQuery(q.query.Deleted())
}

func (q *QueryWrap) Distinct() Query {
	return NewQuery(q.query.Distinct())
}

func (q *QueryWrap) DistinctOn(expr string, params ...interface{}) Query {
	return NewQuery(q.query.DistinctOn(expr, params...))
}

func (q *QueryWrap) DropTable(opt *pgorm.DropTableOptions) error {
	return q.query.DropTable(opt)
}

func (q *QueryWrap) Except(other Query) Query {
	return NewQuery(q.query.Except(other.(*QueryWrap).query))
}

func (q *QueryWrap) ExceptAll(other Query) Query {
	return NewQuery(q.query.ExceptAll(other.(*QueryWrap).query))
}

func (q *QueryWrap) ExcludeColumn(columns ...string) Query {
	return NewQuery(q.query.ExcludeColumn(columns...))
}

func (q *QueryWrap) Exec(query interface{}, params ...interface{}) (pgorm.Result, error) {
	return q.query.Exec(query, params...)
}

func (q *QueryWrap) ExecOne(query interface{}, params ...interface{}) (pgorm.Result, error) {
	return q.query.ExecOne(query, params...)
}

func (q *QueryWrap) Exists() (bool, error) {
	return q.query.Exists()
}

func (q *QueryWrap) First() error {
	return q.query.First()
}

func (q *QueryWrap) For(s string, params ...interface{}) Query {
	return NewQuery(q.query.For(s, params...))
}

func (q *QueryWrap) ForEach(fn interface{}) error {
	return q.query.ForEach(fn)
}

func (q *QueryWrap) ForceDelete(values ...interface{}) (pgorm.Result, error) {
	return q.query.ForceDelete(values...)
}

func (q *QueryWrap) Group(columns ...string) Query {
	return NewQuery(q.query.Group(columns...))
}

func (q *QueryWrap) GroupExpr(group string, params ...interface{}) Query {
	return NewQuery(q.query.GroupExpr(group, params...))
}

func (q *QueryWrap) Having(having string, params ...interface{}) Query {
	return NewQuery(q.query.Having(having, params))
}

func (q *QueryWrap) Insert(values ...interface{}) (pgorm.Result, error) {
	return q.query.Insert(values...)
}

func (q *QueryWrap) Intersect(other Query) Query {
	return NewQuery(q.query.Intersect(other.(*QueryWrap).query))
}

func (q *QueryWrap) IntersectAll(other Query) Query {
	return NewQuery(q.query.IntersectAll(other.(*QueryWrap).query))
}

func (q *QueryWrap) Join(join string, params ...interface{}) Query {
	return NewQuery(q.query.Join(join, params...))
}

func (q *QueryWrap) JoinOn(condition string, params ...interface{}) Query {
	return NewQuery(q.query.JoinOn(condition, params...))
}

func (q *QueryWrap) JoinOnOr(condition string, params ...interface{}) Query {
	return NewQuery(q.query.JoinOnOr(condition, params...))
}

func (q *QueryWrap) Last() error {
	return q.query.Last()
}

func (q *QueryWrap) Limit(n int) Query {
	return NewQuery(q.query.Limit(n))
}

func (q *QueryWrap) Model(model ...interface{}) Query {
	return NewQuery(q.query.Model(model...))
}

func (q *QueryWrap) New() Query {
	return NewQuery(q.query.New())
}

func (q *QueryWrap) Offset(n int) Query {
	return NewQuery(q.query.Offset(n))
}

func (q *QueryWrap) OnConflict(s string, params ...interface{}) Query {
	return NewQuery(q.query.OnConflict(s, params...))
}

func (q *QueryWrap) Order(orders ...string) Query {
	return NewQuery(q.query.Order(orders...))
}

func (q *QueryWrap) OrderExpr(order string, params ...interface{}) Query {
	return NewQuery(q.query.OrderExpr(order, params...))
}

func (q *QueryWrap) Query(model, query interface{}, params ...interface{}) (pgorm.Result, error) {
	return q.query.Query(model, query, params...)
}

func (q *QueryWrap) QueryOne(model, query interface{}, params ...interface{}) (pgorm.Result, error) {
	return q.query.QueryOne(model, query, params...)
}

func (q *QueryWrap) Relation(name string, apply ...func(Query) (Query, error)) Query {
	qfn := make([]func(*pgorm.Query) (*pgorm.Query, error), len(apply))
	for i, fn := range apply {
		newq, err := fn(q)
		qfn[i] = func(pgquery *pgorm.Query) (*pgorm.Query, error) {
			return newq.(*QueryWrap).query, err
		}
	}
	return NewQuery(q.query.Relation(name, qfn...))
}

func (q *QueryWrap) Returning(s string, params ...interface{}) Query {
	return NewQuery(q.query.Returning(s, params...))
}

func (q *QueryWrap) Select(values ...interface{}) error {
	return q.query.Select(values...)
}

func (q *QueryWrap) SelectAndCount(values ...interface{}) (count int, firstErr error) {
	return q.query.SelectAndCount(values...)
}

func (q *QueryWrap) SelectAndCountEstimate(threshold int, values ...interface{}) (count int, firstErr error) {
	return q.query.SelectAndCountEstimate(threshold, values...)
}

func (q *QueryWrap) SelectOrInsert(values ...interface{}) (inserted bool, _ error) {
	return q.query.SelectOrInsert(values...)
}

func (q *QueryWrap) Set(set string, params ...interface{}) Query {
	return NewQuery(q.query.Set(set, params))
}

func (q *QueryWrap) Table(tables ...string) Query {
	return NewQuery(q.query.Table(tables...))
}

func (q *QueryWrap) TableExpr(expr string, params ...interface{}) Query {
	return NewQuery(q.query.TableExpr(expr, params))
}

func (q *QueryWrap) TableModel() pgorm.TableModel {
	return q.query.TableModel()
}

func (q *QueryWrap) Union(other Query) Query {
	return NewQuery(q.query.Union(other.(*QueryWrap).query))
}

func (q *QueryWrap) UnionAll(other Query) Query {
	return NewQuery(q.query.UnionAll(other.(*QueryWrap).query))
}

func (q *QueryWrap) Update(scan ...interface{}) (pgorm.Result, error) {
	return q.query.Update(scan...)
}

func (q *QueryWrap) UpdateNotZero(scan ...interface{}) (pgorm.Result, error) {
	return q.query.UpdateNotZero(scan...)
}

func (q *QueryWrap) Value(column string, value string, params ...interface{}) Query {
	return NewQuery(q.query.Value(column, value, params...))
}

func (q *QueryWrap) Where(condition string, params ...interface{}) Query {
	return NewQuery(q.query.Where(condition, params...))
}

func (q *QueryWrap) WhereGroup(fn func(Query) (Query, error)) Query {
	newq, err := fn(q)
	return NewQuery(q.query.WhereGroup(func(pgquery *pgorm.Query) (*pgorm.Query, error) {
		return newq.(*QueryWrap).query, err
	}))
}

func (q *QueryWrap) WhereIn(where string, slice interface{}) Query {
	return NewQuery(q.query.WhereIn(where, slice))
}

func (q *QueryWrap) WhereInMulti(where string, values ...interface{}) Query {
	return NewQuery(q.query.WhereInMulti(where, values...))
}

func (q *QueryWrap) WhereNotGroup(fn func(Query) (Query, error)) Query {
	newq, err := fn(q)
	return NewQuery(q.query.WhereNotGroup(func(pgquery *pgorm.Query) (*pgorm.Query, error) {
		return newq.(*QueryWrap).query, err
	}))
}

func (q *QueryWrap) WhereOr(condition string, params ...interface{}) Query {
	return NewQuery(q.query.WhereOr(condition, params))
}

func (q *QueryWrap) WhereOrGroup(fn func(Query) (Query, error)) Query {
	newq, err := fn(q)
	return NewQuery(q.query.WhereOrGroup(func(pgquery *pgorm.Query) (*pgorm.Query, error) {
		return newq.(*QueryWrap).query, err
	}))
}

func (q *QueryWrap) WhereOrNotGroup(fn func(Query) (Query, error)) Query {
	newq, err := fn(q)
	return NewQuery(q.query.WhereOrNotGroup(func(pgquery *pgorm.Query) (*pgorm.Query, error) {
		return newq.(*QueryWrap).query, err
	}))
}

func (q *QueryWrap) WherePK() Query {
	return NewQuery(q.query.WherePK())
}

func (q *QueryWrap) WhereStruct(strct interface{}) Query {
	return NewQuery(q.query.WhereStruct(strct))
}

func (q *QueryWrap) With(name string, subq Query) Query {
	return NewQuery(q.query.With(name, subq.(*QueryWrap).query))
}

func (q *QueryWrap) WithDelete(name string, subq Query) Query {
	return NewQuery(q.query.WithDelete(name, subq.(*QueryWrap).query))
}

func (q *QueryWrap) WithInsert(name string, subq Query) Query {
	return NewQuery(q.query.WithInsert(name, subq.(*QueryWrap).query))
}

func (q *QueryWrap) WithUpdate(name string, subq Query) Query {
	return NewQuery(q.query.WithUpdate(name, subq.(*QueryWrap).query))
}

func (q *QueryWrap) WrapWith(name string) Query {
	return NewQuery(q.query.WrapWith(name))
}
