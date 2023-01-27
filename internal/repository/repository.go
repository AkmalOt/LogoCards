package repository

import (
	"LogoForCardsGin/internal/db"
	logging "LogoForCardsGin/logger"
	"LogoForCardsGin/models"
	"context"
	"gorm.io/gorm"
)

type Repository struct {
	Connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{Connection: conn}
}

var Logger = logging.GetLogger()

func (r *Repository) GetUser(context *context.Context) ([]*models.UserCards, error) {
	var users []*models.UserCards
	tx := db.DataB.Table("user_cards").Scan(&users)
	if tx.Error != nil {
		Logger.Error(tx.Error)
		return users, tx.Error
	}

	return users, nil
}

func (r *Repository) AddUser(context *context.Context, user *models.UserCards) error {

	tx := db.DataB.Create(&user)
	if tx.Error != nil {
		Logger.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *Repository) UpdateLogoJson(context *context.Context, userData *models.UserCards) error {

	tx := db.DataB.Model(&models.UserCards{}).Where("id = ?", userData.ID).Updates(models.UserCards{Logo: userData.Logo})
	if tx.Error != nil {
		Logger.Println(tx.Error)
		return tx.Error
		//return fmt.Errorf("ds %w", err)
	}
	return nil
}

//func (r *Repository) UpdateLogoMultipart(context *context.Context, userData *models.UserCards) error {
//	tx := db.DataB.Model(&models.UserCards{}).Where("id = ?", userData.ID).Updates(models.UserCards{Logo: userData.Logo})
//	if tx.Error != nil {
//		log.Println(tx.Error)
//		return tx.Error
//	}
//	return nil
//}

func (r *Repository) ChangeStatus(context *context.Context, userData *models.UserCards) error {

	tx := db.DataB.Model(&userData).Select("active").Updates(models.UserCards{Active: userData.Active})
	if tx.Error != nil {
		Logger.Println(tx.Error)
		return tx.Error
	}
	return nil
}
