package service

import "github.com/musishere/Mustafa-Ecommerce-App/internal/domain"

type UserService struct{}

func (service UserService) Register(input interface{}) (string, error) {
	return "", nil
}
func (service UserService) findUserByEmail(email string) (*domain.User, error) {
	//Perform some db operation and bussiness logic
	return nil, nil
}
func (service UserService) Login(input interface{}) (string, error) {
	return "", nil
}
func (service UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}
func (service UserService) VerifyCode(id uint, code int) error {
	return nil
}
func (service UserService) CreateProfile(id uint, input interface{}) error {
	return nil
}
func (service UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}
func (service UserService) UpdateProfile(id uint, input interface{}) error {
	return nil
}
func (service UserService) BecomeSeller(id uint, input interface{}) (string, error) {
	return "", nil
}
func (service UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}
func (service UserService) CreateCart(input interface{}, user *domain.User) ([]interface{}, error) {
	return nil, nil
}
func (service UserService) CreateOrder(input interface{}, user *domain.User) (int, error) {
	return 0, nil
}
func (service UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	return nil, nil
}
