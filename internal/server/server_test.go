package server

import (
	"testing"

	"gotu/bookstore/internal/handler"

	"github.com/stretchr/testify/assert"
)

// Test for NewServer
func TestNewServer(t *testing.T) {
	apiHandler := &handler.ApiHandler{}
	server := NewServer(apiHandler)

	assert.NotNil(t, server)
	assert.NotNil(t, server.router)
}
