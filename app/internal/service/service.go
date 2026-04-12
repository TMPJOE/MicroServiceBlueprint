package service

import (
	"log/slog"

	"hotel.com/app/internal/repo"
)

type Service interface {
	Check() error
}

type fooService struct {
	l *slog.Logger
	r *repo.Repo
}

func (s *fooService) Check() error {
	s.l.Info("Pinging db...")
	err := s.r.DbPing()
	s.l.Info("is service working", "err", err.Error())
	return err
}

func New(l *slog.Logger, r *repo.Repo) Service {
	return &fooService{
		l: l,
		r: r,
	}
}
