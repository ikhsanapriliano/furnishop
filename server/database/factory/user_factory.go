package factory

import (
	"furnishop/server/enum"
	"furnishop/server/model"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFactory struct {
	db    *gorm.DB
	count int
	role  enum.Role
}

func (u *UserFactory) Seed() {
	for i := 0; i < u.count; i++ {
		data := model.User{
			Id:       uuid.NewString(),
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: faker.Password(),
			Role:     u.role,
			Photo:    faker.FirstNameFemale() + ".jpg",
		}

		err := u.db.Save(&data).Error
		if err != nil {
			panic(err)
		}
	}
}

func NewUserFactory(db *gorm.DB, count int, role enum.Role) *UserFactory {
	return &UserFactory{
		db:    db,
		count: count,
		role:  role,
	}
}
