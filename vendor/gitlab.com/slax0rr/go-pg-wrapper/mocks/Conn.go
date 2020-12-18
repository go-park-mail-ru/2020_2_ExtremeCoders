// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import go_pg_wrapperorm "gitlab.com/slax0rr/go-pg-wrapper/orm"
import io "io"
import mock "github.com/stretchr/testify/mock"
import orm "github.com/go-pg/pg/v9/orm"
import pg "github.com/go-pg/pg/v9"
import pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
import time "time"

// Conn is an autogenerated mock type for the Conn type
type Conn struct {
	mock.Mock
}

// AddQueryHook provides a mock function with given fields: hook
func (_m *Conn) AddQueryHook(hook pg.QueryHook) {
	_m.Called(hook)
}

// Begin provides a mock function with given fields:
func (_m *Conn) Begin() (pgwrapper.Tx, error) {
	ret := _m.Called()

	var r0 pgwrapper.Tx
	if rf, ok := ret.Get(0).(func() pgwrapper.Tx); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgwrapper.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Conn) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *Conn) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// CopyFrom provides a mock function with given fields: r, query, params
func (_m *Conn) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, r, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(io.Reader, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(r, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader, interface{}, ...interface{}) error); ok {
		r1 = rf(r, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyTo provides a mock function with given fields: w, query, params
func (_m *Conn) CopyTo(w io.Writer, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, w, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(io.Writer, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(w, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Writer, interface{}, ...interface{}) error); ok {
		r1 = rf(w, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateComposite provides a mock function with given fields: model, opt
func (_m *Conn) CreateComposite(model interface{}, opt *orm.CreateCompositeOptions) error {
	ret := _m.Called(model, opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, *orm.CreateCompositeOptions) error); ok {
		r0 = rf(model, opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTable provides a mock function with given fields: model, opt
func (_m *Conn) CreateTable(model interface{}, opt *orm.CreateTableOptions) error {
	ret := _m.Called(model, opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, *orm.CreateTableOptions) error); ok {
		r0 = rf(model, opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: model
func (_m *Conn) Delete(model interface{}) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropComposite provides a mock function with given fields: model, opt
func (_m *Conn) DropComposite(model interface{}, opt *orm.DropCompositeOptions) error {
	ret := _m.Called(model, opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, *orm.DropCompositeOptions) error); ok {
		r0 = rf(model, opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropTable provides a mock function with given fields: model, opt
func (_m *Conn) DropTable(model interface{}, opt *orm.DropTableOptions) error {
	ret := _m.Called(model, opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, *orm.DropTableOptions) error); ok {
		r0 = rf(model, opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec provides a mock function with given fields: query, params
func (_m *Conn) Exec(query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...interface{}) error); ok {
		r1 = rf(query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecContext provides a mock function with given fields: c, query, params
func (_m *Conn) ExecContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, c, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(c, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...interface{}) error); ok {
		r1 = rf(c, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecOne provides a mock function with given fields: query, params
func (_m *Conn) ExecOne(query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...interface{}) error); ok {
		r1 = rf(query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecOneContext provides a mock function with given fields: c, query, params
func (_m *Conn) ExecOneContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, c, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(c, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...interface{}) error); ok {
		r1 = rf(c, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForceDelete provides a mock function with given fields: model
func (_m *Conn) ForceDelete(model interface{}) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Formatter provides a mock function with given fields:
func (_m *Conn) Formatter() orm.QueryFormatter {
	ret := _m.Called()

	var r0 orm.QueryFormatter
	if rf, ok := ret.Get(0).(func() orm.QueryFormatter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QueryFormatter)
		}
	}

	return r0
}

// Insert provides a mock function with given fields: model
func (_m *Conn) Insert(model ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, model...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(model...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Model provides a mock function with given fields: model
func (_m *Conn) Model(model ...interface{}) go_pg_wrapperorm.Query {
	var _ca []interface{}
	_ca = append(_ca, model...)
	ret := _m.Called(_ca...)

	var r0 go_pg_wrapperorm.Query
	if rf, ok := ret.Get(0).(func(...interface{}) go_pg_wrapperorm.Query); ok {
		r0 = rf(model...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(go_pg_wrapperorm.Query)
		}
	}

	return r0
}

// ModelContext provides a mock function with given fields: c, model
func (_m *Conn) ModelContext(c context.Context, model ...interface{}) go_pg_wrapperorm.Query {
	var _ca []interface{}
	_ca = append(_ca, c)
	_ca = append(_ca, model...)
	ret := _m.Called(_ca...)

	var r0 go_pg_wrapperorm.Query
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) go_pg_wrapperorm.Query); ok {
		r0 = rf(c, model...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(go_pg_wrapperorm.Query)
		}
	}

	return r0
}

// Param provides a mock function with given fields: param
func (_m *Conn) Param(param string) interface{} {
	ret := _m.Called(param)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// PoolStats provides a mock function with given fields:
func (_m *Conn) PoolStats() *pg.PoolStats {
	ret := _m.Called()

	var r0 *pg.PoolStats
	if rf, ok := ret.Get(0).(func() *pg.PoolStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pg.PoolStats)
		}
	}

	return r0
}

// Prepare provides a mock function with given fields: q
func (_m *Conn) Prepare(q string) (pgwrapper.Stmt, error) {
	ret := _m.Called(q)

	var r0 pgwrapper.Stmt
	if rf, ok := ret.Get(0).(func(string) pgwrapper.Stmt); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgwrapper.Stmt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: model, query, params
func (_m *Conn) Query(model interface{}, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, model, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(model, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(model, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryContext provides a mock function with given fields: c, model, query, params
func (_m *Conn) QueryContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, c, model, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(c, model, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(c, model, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryOne provides a mock function with given fields: model, query, params
func (_m *Conn) QueryOne(model interface{}, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, model, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(model, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(model, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryOneContext provides a mock function with given fields: c, model, query, params
func (_m *Conn) QueryOneContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, c, model, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(c, model, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(c, model, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunInTransaction provides a mock function with given fields: fn
func (_m *Conn) RunInTransaction(fn func(pgwrapper.Tx) error) error {
	ret := _m.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(pgwrapper.Tx) error) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Select provides a mock function with given fields: model
func (_m *Conn) Select(model interface{}) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: model
func (_m *Conn) Update(model interface{}) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithContext provides a mock function with given fields: ctx
func (_m *Conn) WithContext(ctx context.Context) pgwrapper.Conn {
	ret := _m.Called(ctx)

	var r0 pgwrapper.Conn
	if rf, ok := ret.Get(0).(func(context.Context) pgwrapper.Conn); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgwrapper.Conn)
		}
	}

	return r0
}

// WithParam provides a mock function with given fields: param, value
func (_m *Conn) WithParam(param string, value interface{}) pgwrapper.Conn {
	ret := _m.Called(param, value)

	var r0 pgwrapper.Conn
	if rf, ok := ret.Get(0).(func(string, interface{}) pgwrapper.Conn); ok {
		r0 = rf(param, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgwrapper.Conn)
		}
	}

	return r0
}

// WithTimeout provides a mock function with given fields: d
func (_m *Conn) WithTimeout(d time.Duration) pgwrapper.Conn {
	ret := _m.Called(d)

	var r0 pgwrapper.Conn
	if rf, ok := ret.Get(0).(func(time.Duration) pgwrapper.Conn); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgwrapper.Conn)
		}
	}

	return r0
}
