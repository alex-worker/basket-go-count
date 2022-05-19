package context

import (
	"basket-go-count/src/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContext(t *testing.T) {
	fakeConnString := mocks.GetConnectionString()

	ctx, err := NewContext(fakeConnString)
	assert.Nil(t, ctx)
	assert.NotNil(t, err)

	ctx, err = NewContext(mocks.GetConnectionStringRealDb())
	assert.Nil(t, err)
	assert.NotNil(t, ctx)

	ctx.Close()
}

func TestContext_Close(t *testing.T) {
	ctx := Context{db: nil}
	ctx.Close() // we should not fail if db is nil

	ctx = Context{db: mocks.Connect()}
	ctx.Close()
	assert.Nil(t, ctx.db)

	ctx = Context{db: mocks.Connect()}
	err := ctx.db.Ping()
	if err != nil {
		t.Error(err)
	}
	ctx.Close()
	assert.Nil(t, ctx.db)
}

func TestContext_GetDb(t *testing.T) {
	ctx := Context{db: nil}
	actual := ctx.GetDb()
	assert.Nil(t, actual)
}
