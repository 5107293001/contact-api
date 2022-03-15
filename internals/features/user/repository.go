package user

import "gorm.io/gorm"

type RepositoryInterface interface {
	Query(offset, limit int, q string) ([]User, error)
	Get(id uint) (User, error)
	Create(req *User) error
	Update(id uint, update *User) error
	Delete(id uint) error
}

type repository struct {
	db gorm.DB
}

func NewRepository(db gorm.DB) RepositoryInterface {
	return &repository{db}
}

func (repository *repository) Query(offset, limit int, q string) ([]User, error) {
	var datalist [] User
	err := repository.db.Debug().Model(&User{}).
	Where ("name like ? or email like ?", "%"+q+ "%","%"+q+"%").
	Limit(limit).Offset(offset).
	Find (&datalist).Error
	return datalist,err
}

func (repository *repository) Get(id uint) (User, error) {

	user := User{}
	err := repository.db.Debug().Model(&User{}).First(&user,id).Error
	return user, err
}

func (repository *repository) Create(req *User) error {
	return repository.db.Debug().Model(&User{}).Create(&req).Error
}

func (repository *repository) Update(id uint, update *User) error {
	err := repository.db.Debug().Model(&User{}).
	Where("id = ?", id).
	Updates(&update).Error
	return err
}
	


func (repository *repository) Delete(id uint) error {
	err := repository.db.Debug().Model(&User{}).Delete(&User{},id).Error
return err
}
