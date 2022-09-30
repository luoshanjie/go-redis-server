package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-redis-server/internal/domain/mocks"
)

func TestServiceRun(t *testing.T) {
	asserts := assert.New(t)
	mockTCP := new(mocks.TCPDelivery)
	mockTCP.On("ListenAndServer", mock.Anything).Return(nil)
	server := NewBaseService(mockTCP)
	asserts.NoError(server.Run())
}
