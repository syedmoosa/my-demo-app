package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "net/http/httptest"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)
	a := 1
	assert.Equal(a, 1)
}
