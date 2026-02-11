package service

import (
	"context"
	"testing"

	"github.com/gookit/slog"
	"github.com/kstsm/wb-l4.5/internal/dto"
)

func newTestService() Calculator {
	logger := slog.New()
	return NewService(logger)
}

func benchRequest() dto.AddRequest {
	items := make([]string, 100)
	for i := range items {
		items[i] = "string"
	}
	return dto.AddRequest{Items: items}
}

func BenchmarkConcatenate(b *testing.B) {
	svc := newTestService()
	req := benchRequest()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = svc.Concatenate(context.Background(), req)
	}
}
