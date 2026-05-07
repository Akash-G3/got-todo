package repository

import "github.com/sriniously/gotodo/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
