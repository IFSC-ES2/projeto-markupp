package notes

import (
	"context"
	"errors"
	"time"
)

type Note struct {
	ID        string
	Path      string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository interface {
	Save(ctx context.Context, note Note) error
}

var (
	ErrInvalidPath    = errors.New("path inválido")
	ErrInvalidContent = errors.New("content inválido")
	ErrDuplicatePath  = errors.New("path já existe")
)

type Service struct{}

func NewService(repo Repository, maxContentSize int64) *Service {
	return &Service{}
}

func (s *Service) Create(ctx context.Context, path, content string) (Note, error) {
	return Note{}, errors.New("não implementado")
}
