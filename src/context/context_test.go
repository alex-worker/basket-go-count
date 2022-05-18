package context

import (
	"github.com/stretchr/testify/assert"
	"os"
	"syreclabs.com/go/faker"
	"testing"
)

func TestMain(m *testing.M) {
	// Write code here to run before tests

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}

func TestNewContext_Error(t *testing.T) {
	connString := faker.RandomString(10)
	ctx, err := NewContext(connString)
	assert.Nil(t, ctx)
	assert.NotNil(t, err)
}

func TestContext_Close(t *testing.T) {

}

func TestContext_GetDb(t *testing.T) {

}
