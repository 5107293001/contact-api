package user

import "github.com/5107293001/contact-api/internals/models"

type ServiceInterface interface {
	Query(offset, limit int, q string) ([]User, error)
	Get(id uint) (User, error)
	Create(req *User) (User, error)
	Update(id uint, update *User) (User, error)
	Delete(id uint) error
}

type service struct {
	repo RepositoryInterface
}

type User struct {
	models.User
}

func NewService(repo RepositoryInterface) ServiceInterface {
	return &service{repo}
}

func (service *service) Query(offset, limit int, q string) ([]User, error) {
	return service.repo.Query(offset, limit, q)
}
func (service *service) Get(id uint) (User, error) {
	return service.repo.Get(id)
}
func (service *service) Create(req *User) (User, error) {
	err := service.repo.Create(req)
	return *req, err
}
func (service *service) Update(id uint, update *User) (User, error) {
	err := service.repo.Update(id, update)
	return *update, err
}
func (service *service) Delete(id uint) error {
	return service.repo.Delete(id)

}
