package server

import (
	"LogoForCardsGin/internal/repository"
	"LogoForCardsGin/internal/services"
	logging "LogoForCardsGin/logger"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//var Logger = logging.GetLogger()

type Handler struct {
	Engine   *gin.Engine
	Services *services.Services
	Logger   logging.Logger
}

func NewHandler(engine *gin.Engine, services *services.Services, Logger logging.Logger) *Handler {
	return &Handler{
		Engine:   engine,
		Services: services,
		Logger:   Logger,
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
	context := context.Background()
	//var context context.Context
	user, err := h.Services.GetUsers(context)
	if err != nil {
		h.Logger.Errorf("%s in GetUsers(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
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
	context := context.Background()
	var userData UserCards

	request := ctx.Request
	formValue := request.FormValue("data")
	log.Println(formValue)

	err := json.Unmarshal([]byte(formValue), &userData)
	if err != nil {
		h.Logger.Errorf("%s in AddUser(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	file, err := ctx.FormFile("logo")
	if err != nil {
		h.Logger.Errorf("%s - error in FormFile - logo?", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	//Logger.Println(file.Filename)
	randomise := rand.Intn(111111)
	randomiserString := strconv.Itoa(randomise)

	userData.Logo = "./logotypes/" + randomiserString + file.Filename

	err = ctx.SaveUploadedFile(file, userData.Logo)
	if err != nil {
		h.Logger.Println(err, "error in context.SaveUploadedFile")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = h.Services.AddUser(context, (*repository.UserCards)(&userData))
	if err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"AddUser": "Done"})
	//context.String(200, "Done")
}

func (h *Handler) UpdateLogoJson(ctx *gin.Context) {
	context := context.Background()

	var userData *UserCards

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		h.Logger.Println(err, "test")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Services.UpdateUserLogoJson(context, (*repository.UserCards)(userData))
	if err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"UpdateLogoJson": "Done"})
}

func (h *Handler) UpdateLogoMultipart(ctx *gin.Context) {
	context := context.Background()

	var userData UserCards

	request := ctx.Request
	formValue := request.FormValue("data")
	log.Println(formValue)

	err := json.Unmarshal([]byte(formValue), &userData)
	if err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return

	}

	file, err := ctx.FormFile("logo")
	if err != nil {
		log.Println(err, "error in FormFile - logo?")
		return
	}
	log.Println(file.Filename)

	userData.Logo = "./logotypes/" + file.Filename

	err = h.Services.UpdateUserLogoJson(context, (*repository.UserCards)(&userData))
	if err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = ctx.SaveUploadedFile(file, userData.Logo)
	if err != nil {
		h.Logger.Println(err, "error in context.SaveUploadedFile")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"UpdateLogoMultipart": "Done"})

}

func (h *Handler) ChangeStatus(ctx *gin.Context) {
	context := context.Background()

	var userData UserCards

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Services.ChangeStatus(context, (*repository.UserCards)(&userData))
	if err != nil {
		h.Logger.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	//time.Sleep(100 * time.Millisecond)
	ctx.JSON(http.StatusOK, gin.H{"ChangeStatus": "Done"})
}
