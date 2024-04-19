package handler

import (
	"app/internal/auth"
	"app/internal/config"
	"app/internal/download"
	"app/internal/logger"
	sess "app/internal/session"
	"app/internal/user"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc download.DownloadUseCase, luc logger.LogUseCase, auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewDownloadHandler(cfg, uc, luc, auc, suc, uuc)

	downloadRoute := router.Group("/download")
	{
		downloadRoute.GET("/:file_dir/:file_name", h.DownloadFile)
	}
}
