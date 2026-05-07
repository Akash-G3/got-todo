package service

import (
	"github.com/sriniously/gotodo/internal/lib/job"
	"github.com/sriniously/gotodo/internal/repository"
	"github.com/sriniously/gotodo/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
