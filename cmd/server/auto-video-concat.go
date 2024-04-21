package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/mrumyantsev/video-hosting/internal/config"
	"github.com/mrumyantsev/video-hosting/internal/constants"
	qconsts "github.com/mrumyantsev/video-hosting/internal/constants/query"
	"github.com/mrumyantsev/video-hosting/internal/database"
	"github.com/mrumyantsev/video-hosting/internal/stream/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	defTmpDirPath           = "./tmp"
	permDir                 = 0775
	repeatMins              = 60
	tempfileRemoveDelaySecs = 1
)

type NonCatVideo struct {
	Id             int
	CodeMP         string
	StartDatetime  string
	DurationRecord int
}

type Repo struct {
	cfg *config.Config
}

func autoVideoConcat(cfg *config.Config) {
	repo := &Repo{cfg: cfg}

	nonCatPaths, err := repo.getNonconcatedPaths()
	if err != nil {
		log.Fatal("cannot get non concated paths. error:", err)
	}

	if nonCatPaths == nil {
		log.Fatal("no video to concat, restart loop")
	}

	if !usecase.IsPathExists(defTmpDirPath) {
		if err = os.MkdirAll(defTmpDirPath, permDir); err != nil {
			log.Fatal("cannot make a video folder")
		}
	}

	for _, val := range *nonCatPaths {
		paths, err := repo.getVideoPaths(val.CodeMP, val.StartDatetime, val.DurationRecord)
		if err != nil {
			log.Fatal("cannot get video paths from db. error:", err)
		}

		tmpFilePath := defTmpDirPath + "/" + val.CodeMP + ".txt"

		if err = createFile(tmpFilePath); err != nil {
			log.Fatal("cannot create a video")
		}

		if err = fillPathsFile(tmpFilePath, paths); err != nil {
			log.Fatal("cannot fill paths-file properly. error:", err)
		}

		outputVideoPath := tmpFilePath + "/" + fmt.Sprintf("%d_%s.mp4", val.Id, val.CodeMP)

		if err = concatVideo(tmpFilePath, outputVideoPath); err != nil {
			log.Fatal("cannot concat: error in command or output video exists")
		}

		time.Sleep(time.Duration(tempfileRemoveDelaySecs) * time.Second)

		if err = os.Remove(tmpFilePath); err != nil {
			log.Fatal("cannot delete file. error:", err)
		}

		time.Sleep(time.Duration(repeatMins) * time.Minute)
	}
}

func (r *Repo) getNonconcatedPaths() (*[]NonCatVideo, error) {
	r.cfg.DBOName = constants.DBO_L3_Name
	dbo := database.CreateOuterDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, dbo)

	template := qconsts.SELECT_COL_FROM_TBL_WHERE_CND
	col := "\"ID\", \"codeMP\", \"startDatetime\", \"durationRecord\""
	tbl := "\"RequestVideoArchive\""
	cnd := "\"recordStatus\"=1"
	query := fmt.Sprintf(template, col, tbl, cnd)

	rows, err := dbo.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	nonCatPaths := []NonCatVideo{}
	for rows.Next() {
		var nonCatVid NonCatVideo

		if err = rows.Scan(&nonCatVid.Id, &nonCatVid.CodeMP, &nonCatVid.StartDatetime,
			&nonCatVid.DurationRecord); err != nil {
			return nil, err
		}

		nonCatPaths = append(nonCatPaths, nonCatVid)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(nonCatPaths) == 0 {
		return nil, nil
	}

	return &nonCatPaths, nil
}

func (r *Repo) getVideoPaths(pathStream, startDatetime string, durationRecord int) (*[]string, error) {
	r.cfg.DBOName = constants.DBO_L3_Name
	dbo := database.CreateOuterDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, dbo)

	template := qconsts.SELECT_VIDEO_PATH_BETWEEN
	query := fmt.Sprintf(template, pathStream, pathStream, startDatetime, pathStream, startDatetime, durationRecord)

	rows, err := dbo.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var paths []string
	var path string

	for rows.Next() {
		if err = rows.Scan(&path); err != nil {
			return &paths, err
		}

		paths = append(paths, path)
	}

	return &paths, nil
}

func createFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	return nil
}

func fillPathsFile(filepath string, data *[]string) error {
	f, err := os.OpenFile(filepath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	for _, path := range *data {
		line := "file '" + path + "'\n"

		if _, err = f.Write([]byte(line)); err != nil {
			return err
		}
	}

	return nil
}

func concatVideo(pathsFile, outputVideo string) error {
	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i",
		pathsFile, "-c", "copy", outputVideo)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func InitVideoSending(ctx *gin.Context) {
	url := "http://10.100.100.60:8654/api/idRequest"
	method := "POST"
	vidPath := "./tmp/videoplayback.mp4"

	payload := new(bytes.Buffer)

	writer := multipart.NewWriter(payload)
	defer func() { _ = writer.Close() }()

	err := writer.WriteField("id", "123")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(vidPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()

	part2, err := writer.CreateFormFile("file", filepath.Base(vidPath))
	if err != nil {
		log.Fatal(err)
	}

	if _, err = io.Copy(part2, file); err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = res.Body.Close() }()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("accept status: %s\n", string(data))
}
