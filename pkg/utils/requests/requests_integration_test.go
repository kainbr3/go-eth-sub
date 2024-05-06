package requests

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	var result any
	err := New(context.TODO(), "POST", "https://httpbin.org/post", &result, nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
