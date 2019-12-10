package repository

import (
	"log"

	"github.com/samuelmjn/go-library/utils"

	"github.com/jinzhu/gorm"
	"github.com/samuelmjn/go-library/repository/model"
)

// UserRepository :nodoc:
type UserRepository interface {
	Create(req *model.User) (err error)
	FindByID(id int64) (user *model.User, err error)
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepository :nodoc:
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(req *model.User) (err error) {
	req.ID = utils.GenerateID()
	tx := r.db.Begin()
	err = tx.Create(&req).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *userRepo) FindByID(id int64) (user *model.User, err error) {
	var res model.User
	err = r.db.Where("id = ?", id).Take(&res).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	user = &res
	return
}
