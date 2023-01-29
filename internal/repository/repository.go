package repository

import (
	"LogoForCardsGin/internal/db"
	logging "LogoForCardsGin/logger"
	"context"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	Connection *gorm.DB
	Logger     logging.Logger
}

func NewRepository(conn *gorm.DB, Logger logging.Logger) *Repository {
	return &Repository{Connection: conn, Logger: Logger}
}

//var Logger = logging.GetLogger()

func (r *Repository) GetUser(ctx context.Context) ([]*UserCards, error) {
	var users []*UserCards
	tx := db.DataB.Table("user_cards").Scan(&users)
	if tx.Error != nil {
		r.Logger.Error(tx.Error)
		return users, tx.Error
	}

	ctxTimeOut, cancelFunc := context.WithTimeout(ctx, time.Duration(200)*time.Millisecond)
	r.Logger.Println(ctxTimeOut, cancelFunc)
	defer func() {
		r.Logger.Println("GetUser complete")
		cancelFunc()
	}()
	return users, nil
}

func (r *Repository) AddUser(ctx context.Context, user *UserCards) error {

	tx := db.DataB.Create(&user)
	if tx.Error != nil {
		r.Logger.Println(tx.Error)
		return tx.Error
	}
	ctxTimeOut, cancelFunc := context.WithTimeout(ctx, time.Duration(200)*time.Millisecond)
	r.Logger.Println(ctxTimeOut, cancelFunc)
	defer func() {
		r.Logger.Println("AddUser complete")
		cancelFunc()
	}()

	return nil
}

func (r *Repository) UpdateLogoJson(ctx context.Context, userData *UserCards) error {

	tx := db.DataB.Model(&UserCards{}).Where("id = ?", userData.ID).Updates(UserCards{Logo: userData.Logo})
	if tx.Error != nil {
		r.Logger.Println(tx.Error)
		return tx.Error
		//return fmt.Errorf("ds %w", err)
	}
	ctxTimeOut, cancelFunc := context.WithTimeout(ctx, time.Duration(200)*time.Millisecond)
	r.Logger.Println(ctxTimeOut, cancelFunc)
	defer func() {
		r.Logger.Println("UpdateLogoJson complete")
		cancelFunc()
	}()
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

func (r *Repository) ChangeStatus(ctx context.Context, userData *UserCards) error {

	tx := db.DataB.Model(&userData).Select("active").Updates(UserCards{Active: userData.Active})
	if tx.Error != nil {
		r.Logger.Println(tx.Error)
		return tx.Error
	}
	ctxTimeOut, cancelFunc := context.WithTimeout(ctx, time.Duration(200)*time.Millisecond)
	r.Logger.Println(ctxTimeOut, cancelFunc)
	defer func() {
		r.Logger.Println("ChangeStatus complete")
		cancelFunc()
	}()
	return nil
}
