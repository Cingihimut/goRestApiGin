package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService UserService
	ctx         *gin.Context
}

func NewUserController(userService UserService, ctx *gin.Context) UserController {
	return UserController{userService: userService, ctx: ctx}
}

func (uc *UserController) Index(ctx *gin.Context) {
	data := uc.userService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (uc *UserController) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := uc.userService.GetById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (uc *UserController) Create(ctx *gin.Context) {
	data, err := uc.userService.Create(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})

		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (uc *UserController) Delete(ctx *gin.Context) {
	data, err := uc.userService.Delete(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})

		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (uc *UserController) Update(ctx *gin.Context) {
	data, err := uc.userService.Update(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})

		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}
