package server

import (
	"LogoForCardsGin/internal/services"
	logging "LogoForCardsGin/logger"
	"LogoForCardsGin/models"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var Logger = logging.GetLogger()

type Handler struct {
	Engine   *gin.Engine
	Services *services.Services
	//Logger *logging.Logger
}

func NewHandler(engine *gin.Engine, services *services.Services) *Handler {
	return &Handler{
		Engine:   engine,
		Services: services,
		//Logger: logger,
	}
}

func (h *Handler) Init() {
	h.Engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Connected"})
	})

	h.Engine.GET("/get_users", h.GetUsers)
	h.Engine.POST("/add_user", h.AddUser)
	h.Engine.POST("/update_logo", h.UpdateLogoJson)
	h.Engine.POST("/update_logo_multi", h.UpdateLogoMultipart)
	h.Engine.POST("/change_status", h.ChangeStatus)
}

func (h *Handler) GetUsers(ctx *gin.Context) {
	//var users []*models.UserCards
	
	var context *context.Context
	user, err := h.Services.GetUsers(context)
	if err != nil {
		Logger.Errorf("%s in GetUsers(server)", err)

		return
	}
	//w := cfx.Writer
	//for _, test := range user {
	//	file, err := os.OpenFile(test.Logo, os.O_CREATE|os.O_RDWR, 0777)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	defer func(file *os.File) {
	//		err := file.Close()
	//		if err != nil {
	//			log.Println(err)
	//		}
	//	}(file)
	//
	//	_, err = io.Copy(w, file)
	//	if err != nil {
	//		return
	//	}
	//	context.JSON(http.StatusOK, test)
	//}
	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) AddUser(ctx *gin.Context) {
	var context *context.Context
	var userData models.UserCards

	request := ctx.Request
	formValue := request.FormValue("data")
	log.Println(formValue)

	err := json.Unmarshal([]byte(formValue), &userData)
	if err != nil {
		Logger.Errorf("%s in AddUser(server)", err)
		return
	}
	file, err := ctx.FormFile("logo")
	if err != nil {
		Logger.Errorf("%s - error in FormFile - logo?", err)
		return
	}
	//Logger.Println(file.Filename)

	userData.Logo = "./logotypes/" + file.Filename

	err = ctx.SaveUploadedFile(file, userData.Logo)
	if err != nil {
		Logger.Println(err, "error in context.SaveUploadedFile")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = h.Services.AddUser(context, &userData)
	if err != nil {
		Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"AddUser": "Done"})
	//context.String(200, "Done")
}

func (h *Handler) UpdateLogoJson(ctx *gin.Context) {
	var context *context.Context

	var userData *models.UserCards

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		Logger.Println(err, "test")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Services.UpdateUserLogoJson(context, userData)
	if err != nil {
		Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"UpdateLogoJson": "Done"})
}

func (h *Handler) UpdateLogoMultipart(ctx *gin.Context) {
	var context *context.Context

	var userData models.UserCards

	request := ctx.Request
	formValue := request.FormValue("data")
	log.Println(formValue)

	err := json.Unmarshal([]byte(formValue), &userData)
	if err != nil {
		log.Println(err)
		return

	}

	file, err := ctx.FormFile("logo")
	if err != nil {
		log.Println(err, "error in FormFile - logo?")
		return
	}
	log.Println(file.Filename)

	userData.Logo = "./logotypes/" + file.Filename

	err = h.Services.UpdateUserLogoJson(context, &userData)
	if err != nil {
		Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = ctx.SaveUploadedFile(file, userData.Logo)
	if err != nil {
		Logger.Println(err, "error in context.SaveUploadedFile")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"UpdateLogoMultipart": "Done"})

}

func (h *Handler) ChangeStatus(ctx *gin.Context) {
	var context *context.Context

	var userData models.UserCards

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Services.ChangeStatus(context, &userData)
	if err != nil {
		Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	//time.Sleep(100 * time.Millisecond)
	ctx.JSON(http.StatusOK, gin.H{"ChangeStatus": "Done"})
}
