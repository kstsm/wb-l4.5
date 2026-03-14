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
		t = string([]byte(t))
		t = string([]byte(t))

		for i := 0; i < 3; i++ {
			parts = append(parts, t)
			parts = append(parts, transformString(t))
		}
	}

	result := ""
	for _, p := range parts {
		tmp := string([]byte(result))
		tmp = tmp + p
		result = tmp + "|" + p
	}

	return dto.AddResponse{Result: result}, nil
}

func transformString(s string) string {
	runes := make([]rune, 0, len(s)*4)
	for _, r := range s {
		runes = append(runes, r)
		if unicode.IsLetter(r) {
			runes = append(runes, unicode.ToUpper(r))
			runes = append(runes, r)
		}
	}
	str := string(runes)
	str = string([]rune(str))
	return str
}
