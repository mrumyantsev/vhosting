package handler

import (
	"github.com/mrumyantsev/video-hosting/internal/auth"
	"github.com/mrumyantsev/video-hosting/internal/config"
	sconfig "github.com/mrumyantsev/video-hosting/internal/config-stream"
	"github.com/mrumyantsev/video-hosting/internal/info"
	"github.com/mrumyantsev/video-hosting/internal/logger"
	sess "github.com/mrumyantsev/video-hosting/internal/session"
	"github.com/mrumyantsev/video-hosting/internal/user"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, scfg *sconfig.Config, uc info.InfoUseCase, luc logger.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewInfoHandler(cfg, scfg, uc, luc, auc, suc, uuc)

	infoRoute := router.Group("/info")
	{
		infoRoute.POST("", h.CreateInfo)
		infoRoute.GET(":id", h.GetInfo)
		infoRoute.GET("all", h.GetAllInfos)
		infoRoute.PATCH(":id", h.PartiallyUpdateInfo)
		infoRoute.DELETE(":id", h.DeleteInfo)
	}
}
