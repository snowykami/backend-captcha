package api

import (
	"git.liteyuki.icu/backend/golite/hertz"
	"git.liteyuki.icu/backend/golite/hertz/middleware"
	"git.liteyuki.icu/backend/golite/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

var h *server.Hertz

func init() {
	logger.Info("Init Router...")
	h = hertz.NewHertz(nil, []app.HandlerFunc{middleware.Cors})
	verify := h.Group("/verify")
	{
		verify.POST("/create", verifyCreate)
		verify.POST("/check", verifyCheck)
	}
}

// Run 内部使用goroutine启动
func Run() {
	go func() {
		err := h.Run()
		if err != nil {
		}
	}()
}
