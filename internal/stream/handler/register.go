package handler

import (
	"github.com/mrumyantsev/video-hosting/internal/auth"
	"github.com/mrumyantsev/video-hosting/internal/config"
	sconfig "github.com/mrumyantsev/video-hosting/internal/config-stream"
	"github.com/mrumyantsev/video-hosting/internal/logger"
	sess "github.com/mrumyantsev/video-hosting/internal/session"
	"github.com/mrumyantsev/video-hosting/internal/stream"
	"github.com/mrumyantsev/video-hosting/internal/user"

	"github.com/gin-gonic/gin"
)

func RegisterTemplateHTTPEndpoints(router *gin.Engine, cfg *config.Config, scfg *sconfig.Config, uc stream.StreamUseCase,
	uuc user.UserUseCase, luc logger.LogUseCase, auc auth.AuthUseCase, suc sess.SessUseCase) {
	h := NewStreamHandler(cfg, scfg, uc, uuc, luc, auc, suc)

	router.GET("/stream", h.ServeIndex)
	router.GET("/stream/:uuid", h.ServeStream)
}

func RegisterStreamingHTTPEndpoints(router *gin.Engine, cfg *config.Config, scfg *sconfig.Config, uc stream.StreamUseCase,
	uuc user.UserUseCase, luc logger.LogUseCase, auc auth.AuthUseCase, suc sess.SessUseCase) {
	h := NewStreamHandler(cfg, scfg, uc, uuc, luc, auc, suc)

	streamRoute := router.Group("/stream")
	{
		streamRoute.GET("/codec/:uuid", h.ServeStreamCodec)
		streamRoute.POST("/receiver/:uuid", h.ServeStreamVidOverWebRTC)
		streamRoute.POST("/", h.ServeStreamWebRTC2)

		streamRoute.GET("/get/:id", h.GetStream)
		streamRoute.GET("/get/all", h.GetAllStreams)
	}
}
