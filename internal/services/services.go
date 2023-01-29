package services

import (
	"LogoForCardsGin/internal/repository"
	logging "LogoForCardsGin/logger"
	"context"
	"strconv"
)

type Services struct {
	Repository *repository.Repository
	Logger     logging.Logger
}

func NewServices(rep *repository.Repository, Logger logging.Logger) *Services {
	return &Services{Repository: rep, Logger: Logger}
}

//var Logger = logging.GetLogger()

func (s *Services) GetUsers(context context.Context) ([]*repository.UserCards, error) {
	return s.Repository.GetUser(context)
}

func (s *Services) AddUser(context context.Context, user *repository.UserCards) (err error) {

	length := len(strconv.Itoa(user.Num))
	if length > 20 {
		s.Logger.Println("error - Card length is more than 19")
		return err
	}

	err = s.Repository.AddUser(context, user)
	if err != nil {
		s.Logger.Println(err)
		return err
	}
	return nil
}

func (s *Services) UpdateUserLogoJson(context context.Context, userData *repository.UserCards) error {
	return s.Repository.UpdateLogoJson(context, userData)
}

//func (s *Services) UpdateLogoMultipart(context *context.Context, userData *models.UserCards) error {
//	s.Repository.UpdateLogoMultipart(context, userData)
//	return nil
//}

func (s *Services) ChangeStatus(context context.Context, userData *repository.UserCards) error {
	return s.Repository.ChangeStatus(context, userData)
}
