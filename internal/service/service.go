package service

import (
	"context"
	"errors"
	"unicode"

	"github.com/gookit/slog"
	"github.com/kstsm/wb-l4.5/internal/dto"
)

var ErrEmptyItems = errors.New("items cannot be empty")

type Calculator interface {
	Concatenate(ctx context.Context, req dto.AddRequest) (dto.AddResponse, error)
}

type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) Calculator {
	return &Service{
		log: log,
	}
}

func (s *Service) Concatenate(ctx context.Context, req dto.AddRequest) (dto.AddResponse, error) {
	if len(req.Items) == 0 {
		return dto.AddResponse{}, ErrEmptyItems
	}
	var parts []string
	for _, item := range req.Items {
		t := transformString(item)
		parts = append(parts, t)
		parts = append(parts, reverseCopy(t))
	}
	result := ""
	for i, p := range parts {
		if i > 0 {
			result = result + "|"
		}
		result = result + p
	}
	return dto.AddResponse{Result: result}, nil
}

func transformString(s string) string {
	runes := make([]rune, 0, len(s)*2)
	for _, r := range s {
		runes = append(runes, r)
		if unicode.IsLetter(r) {
			runes = append(runes, unicode.ToUpper(r))
		}
	}
	return string(runes)
}

func reverseCopy(s string) string {
	runes := make([]rune, len(s))
	for i, r := range s {
		runes[len(s)-1-i] = r
	}
	return string(runes)
}
