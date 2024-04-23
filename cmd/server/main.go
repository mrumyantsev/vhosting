package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/video-hosting/internal/server"
	"github.com/mrumyantsev/video-hosting/internal/server/middleware"

	"github.com/joho/godotenv"
	authhandler "github.com/mrumyantsev/video-hosting/internal/auth/handler"
	authrepo "github.com/mrumyantsev/video-hosting/internal/auth/repository"
	authusecase "github.com/mrumyantsev/video-hosting/internal/auth/usecase"
	"github.com/mrumyantsev/video-hosting/internal/config"
	sconfig "github.com/mrumyantsev/video-hosting/internal/config-stream"
	downloadhandler "github.com/mrumyantsev/video-hosting/internal/download/handler"
	downloadusecase "github.com/mrumyantsev/video-hosting/internal/download/usecase"
	grouphandler "github.com/mrumyantsev/video-hosting/internal/group/handler"
	grouprepo "github.com/mrumyantsev/video-hosting/internal/group/repository"
	groupusecase "github.com/mrumyantsev/video-hosting/internal/group/usecase"
	infohandler "github.com/mrumyantsev/video-hosting/internal/info/handler"
	inforepo "github.com/mrumyantsev/video-hosting/internal/info/repository"
	infousecase "github.com/mrumyantsev/video-hosting/internal/info/usecase"
	"github.com/mrumyantsev/video-hosting/internal/logger"
	logrepo "github.com/mrumyantsev/video-hosting/internal/logger/repository"
	logusecase "github.com/mrumyantsev/video-hosting/internal/logger/usecase"
	msg "github.com/mrumyantsev/video-hosting/internal/messages"
	permhandler "github.com/mrumyantsev/video-hosting/internal/permission/handler"
	permrepo "github.com/mrumyantsev/video-hosting/internal/permission/repository"
	permusecase "github.com/mrumyantsev/video-hosting/internal/permission/usecase"
	sessrepo "github.com/mrumyantsev/video-hosting/internal/session/repository"
	sessusecase "github.com/mrumyantsev/video-hosting/internal/session/usecase"
	streamhandler "github.com/mrumyantsev/video-hosting/internal/stream/handler"
	streamrepo "github.com/mrumyantsev/video-hosting/internal/stream/repository"
	streamusecase "github.com/mrumyantsev/video-hosting/internal/stream/usecase"
	userhandler "github.com/mrumyantsev/video-hosting/internal/user/handler"
	userrepo "github.com/mrumyantsev/video-hosting/internal/user/repository"
	userusecase "github.com/mrumyantsev/video-hosting/internal/user/usecase"
	videohandler "github.com/mrumyantsev/video-hosting/internal/video/handler"
	videorepo "github.com/mrumyantsev/video-hosting/internal/video/repository"
	videousecase "github.com/mrumyantsev/video-hosting/internal/video/usecase"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/mrumyantsev/video-hosting/docs"
)

const (
	configsDir string = "./configs/"
)

// @title       Video Hosting
// @version     1.0
// @description Server API for Video Hosting Service

// @host     localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Read .env file
	err := godotenv.Load(configsDir + ".env")
	if err != nil {
		logger.Print(msg.FatalFailedToLoadEnvironmentFile(err))

		return
	}

	logger.Print(msg.InfoEnvironmentsLoaded())

	// Load config
	cfg, err := config.LoadConfig(configsDir + "config.yml")
	if err != nil {
		logger.Print(msg.FatalFailedToLoadConfigFile(err))

		return
	}

	logger.Print(msg.InfoConfigLoaded())

	// Make stream config
	scfg := new(sconfig.Config)

	// Init repositories
	userRepo := userrepo.NewUserRepository(cfg)
	authRepo := authrepo.NewAuthRepository(cfg)
	sessRepo := sessrepo.NewSessRepository(cfg)
	logRepo := logrepo.NewLogRepository(cfg)
	groupRepo := grouprepo.NewGroupRepository(cfg)
	permRepo := permrepo.NewPermRepository(cfg)
	infoRepo := inforepo.NewInfoRepository(cfg)
	videoRepo := videorepo.NewVideoRepository(cfg)
	streamRepo := streamrepo.NewStreamRepository(cfg)

	// Init use cases
	userUseCase := userusecase.NewUserUseCase(cfg, userRepo)
	authUseCase := authusecase.NewAuthUseCase(cfg, authRepo)
	sessUseCase := sessusecase.NewSessUseCase(sessRepo, authRepo)
	logUseCase := logusecase.NewLogUseCase(logRepo)
	groupUseCase := groupusecase.NewGroupUseCase(groupRepo)
	permUseCase := permusecase.NewPermUseCase(permRepo)
	infoUseCase := infousecase.NewInfoUseCase(infoRepo)
	videoUseCase := videousecase.NewVideoUseCase(videoRepo)
	streamUC := streamusecase.NewStreamUseCase(cfg, scfg, streamRepo)
	downloadUseCase := downloadusecase.NewDownloadUseCase(cfg)

	// Init engine
	router := gin.New()

	// Add middleware
	router.Use(middleware.CORSMiddleware())

	// Register static web files for stream
	if _, err := os.Stat("./web"); !os.IsNotExist(err) {
		router.LoadHTMLGlob("./web/templates/*")

		streamhandler.RegisterTemplateHTTPEndpoints(router, cfg, scfg, streamUC,
			userUseCase, logUseCase, authUseCase, sessUseCase)
	}

	router.StaticFS("/static", http.Dir("./web/static"))

	// Register routes
	authhandler.RegisterHTTPEndpoints(router, cfg, authUseCase, userUseCase,
		sessUseCase, logUseCase)
	userhandler.RegisterHTTPEndpoints(router, cfg, userUseCase, logUseCase,
		authUseCase, sessUseCase)
	grouphandler.RegisterHTTPEndpoints(router, cfg, groupUseCase, logUseCase,
		authUseCase, sessUseCase, userUseCase)
	permhandler.RegisterHTTPEndpoints(router, cfg, permUseCase, logUseCase,
		authUseCase, sessUseCase, userUseCase, groupUseCase)
	infohandler.RegisterHTTPEndpoints(router, cfg, scfg, infoUseCase, logUseCase,
		authUseCase, sessUseCase, userUseCase)
	videohandler.RegisterHTTPEndpoints(router, cfg, videoUseCase, logUseCase,
		authUseCase, sessUseCase, userUseCase)
	streamhandler.RegisterStreamingHTTPEndpoints(router, cfg, scfg, streamUC,
		userUseCase, logUseCase, authUseCase, sessUseCase)
	downloadhandler.RegisterHTTPEndpoints(router, cfg, downloadUseCase, logUseCase,
		authUseCase, sessUseCase, userUseCase)

	// Register Swagger handler for Server API documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Debug mode
	if cfg.ServerDebugEnable {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init HTTP server
	srv := server.New(cfg, router)

	isShutdown := false

	// Start HTTP server
	go func() {
		if err := srv.Start(); (err != nil) && !isShutdown {
			logger.Print(msg.FatalFailureOnServerRunning(err))
		}
	}()

	cfg.ServerIP = getOutboundIP()

	logger.Print(msg.InfoServerStartedSuccessfullyAtLocalAddress(cfg.ServerIP, cfg.ServerPort))

	// Start videostreams service
	go streamUC.ServeStreams()

	// Start video gluing routine
	go autoVideoConcat(cfg)

	// Listen for interrupt signal from keyboard
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit

	logger.Print(msg.InfoRecivedSignal(sig))

	isShutdown = true

	// Shut down HTTP server
	ctx, shutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdown()

	if err = srv.Shutdown(ctx); err != nil {
		logger.Print(msg.FatalFailureOnServerShutdown(err))

		return
	}

	logger.Print(msg.InfoServerShutedDownCorrectly())
}

func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		logger.Print(msg.WarningCannotGetLocalIP(err))
	}

	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}
