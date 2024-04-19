package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/video-hosting/internal/auth"
	"github.com/mrumyantsev/video-hosting/internal/config"
	"github.com/mrumyantsev/video-hosting/internal/group"
	"github.com/mrumyantsev/video-hosting/internal/logger"
	perm "github.com/mrumyantsev/video-hosting/internal/permission"
	sess "github.com/mrumyantsev/video-hosting/internal/session"
	"github.com/mrumyantsev/video-hosting/internal/user"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc perm.PermUseCase, luc logger.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase, guc group.GroupUseCase) {
	h := NewPermHandler(cfg, uc, luc, auc, suc, uuc, guc)

	permRoute := router.Group("/perm")
	{
		permRoute.GET("all", h.GetAllPermissions)
	}

	permSetUserRoute := router.Group("/perm/user")
	{
		permSetUserRoute.POST(":id", h.SetUserPermissions)
		permSetUserRoute.GET(":id", h.GetUserPermissions)
		permSetUserRoute.DELETE(":id", h.DeleteUserPermissions)
	}

	permSetGroupRoute := router.Group("/perm/group")
	{
		permSetGroupRoute.POST(":id", h.SetGroupPermissions)
		permSetGroupRoute.GET(":id", h.GetGroupPermissions)
		permSetGroupRoute.DELETE(":id", h.DeleteGroupPermissions)
	}
}
