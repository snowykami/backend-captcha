package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"liteyuki-captcha/service"
)

func verifyCreate(c context.Context, ctx *app.RequestContext) {
	id, b64, _, err := service.GenerateCaptcha()
	if err != nil {
		ctx.JSON(
			consts.StatusInternalServerError,
			utils.H{
				"message": err.Error(),
			},
		)
	}
	ctx.JSON(
		consts.StatusOK,
		utils.H{
			"message": "success",
			"payload": utils.H{
				"id":    id,
				"image": b64,
			},
		})
}

func verifyCheck(c context.Context, ctx *app.RequestContext) {
	id := ctx.PostForm("id")
	ans := ctx.PostForm("answer")
	if service.VerifyCaptcha(id, ans) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "success",
			"payload": utils.H{"result": true}},
		)
	} else {
		ctx.JSON(consts.StatusInternalServerError,
			utils.H{"message": "success",
				"payload": utils.H{"result": false}},
		)
	}
}
