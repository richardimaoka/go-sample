package users

import (
	"fmt"

	"github.com/richardimaoka/go-sandbox/applog"
)

// エンティティの定義

type User struct {
	ID   int
	Name string
}

// リポジトリの定義

type Repository interface {
	FindByID(id int) *User
}

func NewStubRepository(l applog.Logger) Repository {
	return &stubRepository{
		logger: l,
		users: map[int]*User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}
}

var _ Repository = (*stubRepository)(nil)

type stubRepository struct {
	logger applog.Logger
	users  map[int]*User
}

func (r *stubRepository) FindByID(id int) *User {
	user, exists := r.users[id]
	if !exists {
		r.logger.Error("User not found")
		return nil
	}
	r.logger.Info(fmt.Sprintf("Fetched user with ID %d", id))
	return user
}

// サービスの定義

func NewService(logger applog.Logger, repo Repository) *Service {
	return &Service{logger: logger, repo: repo}
}

type Service struct {
	logger applog.Logger
	repo   Repository
}

func (s *Service) GetUserByID(id int) *User {
	s.logger.Info(fmt.Sprintf("Invoking repo.FindByID with %d", id))
	return s.repo.FindByID(id)
}
