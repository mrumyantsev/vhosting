package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/video-hosting/internal/auth"
	"github.com/mrumyantsev/video-hosting/internal/config"
	"github.com/mrumyantsev/video-hosting/internal/group"
	"github.com/mrumyantsev/video-hosting/internal/logger"
	sess "github.com/mrumyantsev/video-hosting/internal/session"
	"github.com/mrumyantsev/video-hosting/internal/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc group.GroupUseCase, luc logger.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewGroupHandler(cfg, uc, luc, auc, suc, uuc)

	groupRoute := router.Group("/group")
	{
		groupRoute.POST("", h.CreateGroup)
		groupRoute.GET(":id", h.GetGroup)
		groupRoute.GET("all", h.GetAllGroups)
		groupRoute.PATCH(":id", h.PartiallyUpdateGroup)
		groupRoute.DELETE(":id", h.DeleteGroup)
	}

	groupSetUserRoute := router.Group("/group/user")
	{
		groupSetUserRoute.POST(":id", h.SetUserGroups)
		groupSetUserRoute.GET(":id", h.GetUserGroups)
		groupSetUserRoute.DELETE(":id", h.DeleteUserGroups)
	}
}
