package handler

import (
	"app/internal/auth"
	"app/internal/config"
	"app/internal/group"
	"app/internal/logger"
	sess "app/internal/session"
	"app/internal/user"
	"github.com/gin-gonic/gin"
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
