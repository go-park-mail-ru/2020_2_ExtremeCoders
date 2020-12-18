# go-pg-wrapper

A wrapper around [go-pg](https://github.com/go-pg/pg) providing functions that
use interfaces, for simpler unit testing.

## Usage

Connect to your database using `go-pg` as you normally would, and pass the
`*pg.DB` to `pgwrapper.NewDB()`:

```go
package main

import (
	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
	"github.com/go-pg/pg/v9"
)

func main() {
	db := pgwrapper.NewDB(pg.Connect(&pg.Options{
		// ...
	}))

    // use db as before
}
```

### Tests

Mocks are also provided implementing the
[testify mock](https://godoc.org/github.com/stretchr/testify/mock) package for
your tests:

```go
package foo

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/slax0rr/go-pg-wrapper/mocks"
)

func TestFoo(t *testing.T) {
	expObj := &FooObj{
		// ...
	}

	db := new(mocks.DB)
	db.On("Model", mock.Anything).Return(db)
	db.On("Select").Return(expObj, nil)

	foo, err := Foo(db)
	assert.Nil(t, err)
	assert.Equal(t, expObj, foo)
}
```
