package handler

import (
	"github.com/mrumyantsev/video-hosting/internal/auth"
	"github.com/mrumyantsev/video-hosting/internal/config"
	"github.com/mrumyantsev/video-hosting/internal/download"
	"github.com/mrumyantsev/video-hosting/internal/logger"
	sess "github.com/mrumyantsev/video-hosting/internal/session"
	"github.com/mrumyantsev/video-hosting/internal/user"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc download.DownloadUseCase, luc logger.LogUseCase, auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewDownloadHandler(cfg, uc, luc, auc, suc, uuc)

	downloadRoute := router.Group("/download")
	{
		downloadRoute.GET("/:file_dir/:file_name", h.DownloadFile)
	}
}
