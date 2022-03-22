package repositories

import (
	"github.com/maoaeri/openapi/pkg/repositories/postrepo"
	"github.com/maoaeri/openapi/pkg/repositories/userrepo"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepo *userrepo.UserRepo
	PostRepo *postrepo.PostRepo
}

func InitRepository(db *gorm.DB) *Repository {
	userRepo := userrepo.NewUserRepo(db)
	postrepo := postrepo.NewPostRepo(db)
	return &Repository{
		UserRepo: userRepo,
		PostRepo: postrepo,
	}
}
