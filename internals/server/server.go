package server

import (
	"github.com/5107293001/contact-api/internals/features/user"
	"github.com/5107293001/contact-api/pkg/db/postgres"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type server struct {
	C  *gin.Engine
	DB *gorm.DB
}

func GetServer() *server {
	return &server{
		C:  gin.Default(),
		DB: postgres.ConnetDatabase(),
	}
}
func (s *server) Run() {
	s.initRoutes()
	s.C.Run()
}
func (s *server) initRoutes() {
	user.RegisterRoutes(s.C, user.NewService(user.NewRepository(*s.DB)))

}
