package database

import "SplitPay_back/internal/domain"

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(u *domain.Wari_user) {
	db.Create(&u)
}

func (db *UserRepository) Select() []domain.Wari_user {
	user := []domain.Wari_user{}
	db.FindAll(&user)
	return user
}
func (db *UserRepository) Delete(id string) {
	user := []domain.Wari_user{}
	db.DeleteById(&user, id)
}
