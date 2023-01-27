package services

import (
	"LogoForCardsGin/internal/repository"
	logging "LogoForCardsGin/logger"
	"LogoForCardsGin/models"
	"context"
	"strconv"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}

var Logger = logging.GetLogger()

func (s *Services) GetUsers(context *context.Context) ([]*models.UserCards, error) {
	return s.Repository.GetUser(context)
}

func (s *Services) AddUser(context *context.Context, user *models.UserCards) (err error) {

	length := len(strconv.Itoa(user.Num))
	if length > 20 {
		Logger.Println("error - Card length is more than 19")
		return err
	}

	err = s.Repository.AddUser(context, user)
	if err != nil {
		Logger.Println(err)
		return err
	}
	return nil
}

func (s *Services) UpdateUserLogoJson(context *context.Context, userData *models.UserCards) error {
	return s.Repository.UpdateLogoJson(context, userData)
}

//func (s *Services) UpdateLogoMultipart(context *context.Context, userData *models.UserCards) error {
//	s.Repository.UpdateLogoMultipart(context, userData)
//	return nil
//}

func (s *Services) ChangeStatus(context *context.Context, userData *models.UserCards) error {
	return s.Repository.ChangeStatus(context, userData)
}
