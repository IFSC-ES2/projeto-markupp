package notes

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
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
	GetNoteByID(ctx context.Context, id string) (Note, error)
}

var (
	ErrInvalidPath    = errors.New("path inválido")
	ErrInvalidContent = errors.New("content inválido")
	ErrDuplicatePath  = errors.New("path já existe")
	ErrInvalidId      = errors.New("ID inválido")
	ErrNotFoundId     = errors.New("ID não encontrado")
)

type Service struct {
	repo           Repository
	clock          func() time.Time
	newID          func() string
	maxContentSize int64
}

func NewService(repo Repository, maxContentSize int64) *Service {
	return &Service{
		repo:           repo,
		clock:          time.Now,
		newID:          func() string { return uuid.NewString() },
		maxContentSize: maxContentSize,
	}
}

func (s *Service) GetNoteById(ctx context.Context, id string) (Note, error) {
	if err := s.validateId(ctx, id); err != nil {
		return Note{}, err
	}

	note, err := s.repo.GetNoteByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return Note{}, ErrNotFoundId
	}
	return note, err
}

func (s *Service) Create(ctx context.Context, path, content string) (Note, error) {
	if err := validatePath(path); err != nil {
		return Note{}, err
	}
	if err := s.validateContent(content); err != nil {
		return Note{}, err
	}
	now := s.clock()
	note := Note{
		ID:        s.newID(),
		Path:      path,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.repo.Save(ctx, note); err != nil {
		return Note{}, fmt.Errorf("salvar nota %q: %w", note.ID, err)
	}
	return note, nil
}

func (s *Service) validateId(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidId
	}
	return nil
}

func validatePath(path string) error {
	p := strings.TrimSpace(path)
	if p == "" {
		return ErrInvalidPath
	}
	if len(p) > 1024 {
		return ErrInvalidPath
	}
	if strings.Contains(p, "..") {
		return ErrInvalidPath
	}
	return nil
}

func (s *Service) validateContent(content string) error {
	if int64(len(content)) > s.maxContentSize {
		return ErrInvalidContent
	}
	return nil
}
