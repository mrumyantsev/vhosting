package repository

import (
	"fmt"

	"github.com/mrumyantsev/video-hosting/internal/config"
	qconsts "github.com/mrumyantsev/video-hosting/internal/constants/query"
	"github.com/mrumyantsev/video-hosting/internal/database"
	"github.com/mrumyantsev/video-hosting/internal/logger"
)

type LogRepository struct {
	cfg *config.Config
}

func NewLogRepository(cfg *config.Config) *LogRepository {
	return &LogRepository{cfg: cfg}
}

func (r *LogRepository) CreateLogRecord(log *logger.Log) error {
	db := database.CreateLocalDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := qconsts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s, %s, %s, %s, %s, %s, %s, %s)", logger.TableName,
		logger.ErrLevel, logger.ClientID, logger.SessionOwner, logger.RequestMethod, logger.RequestPath,
		logger.StatusCode, logger.ErrCode, logger.Message, logger.CreationDate)
	val := "($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	query := fmt.Sprintf(template, tbl, val)

	if _, err := db.Query(query, log.ErrLevel, log.ClientIP, log.SessionOwner,
		log.RequestMethod, log.RequestPath, log.StatusCode, log.ErrCode,
		log.Message.(string), log.CreationDate); err != nil {
		return err
	}

	return nil
}
