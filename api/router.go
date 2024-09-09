package api

import (
	"git.liteyuki.icu/backend/golite/hertz"
	"git.liteyuki.icu/backend/golite/logger"
	"github.com/cloudwego/hertz/pkg/app/server"
)

var h *server.Hertz

func init() {
	logger.Info("Init Router...")
	h = hertz.NewHertz()
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
