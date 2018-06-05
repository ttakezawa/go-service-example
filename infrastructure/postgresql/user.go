package postgresql

import (
	"github.com/jinzhu/gorm"
	"github.com/ttakezawa/go-service-example/domain"
)

// UserRepository provides methods to persist
type UserRepository struct {
	DB *gorm.DB `inject:""`
}

var _ domain.UserRepository = (*UserRepository)(nil)

// FindByName returns User
func (userRepository *UserRepository) FindByName(name string) (*domain.User, error) {
	var user domain.User
	if result := userRepository.DB.Where("name = ?", name).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
