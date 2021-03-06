package rest

import (
	"fmt"
	"net/http"

	"github.com/ditointernet/tradulab-service/drivers"
	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/gin-gonic/gin"
)

type ServiceInput struct {
	File services.FileHandler
}

type File struct {
	in ServiceInput
}

func NewFile(in ServiceInput) (*File, error) {
	if in.File == nil {
		return nil, fmt.Errorf("error message")
	}

	return &File{in: in}, nil
}

func (f File) CreateFile(ctx *gin.Context) {
	body := &drivers.File{}

	err := ctx.ShouldBindJSON(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	file := &domain.File{
		ProjectId: body.ProjectID,
		FileName:  body.FileName,
	}

	if body.ProjectID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "project_id is required",
		})
		return
	}

	newFile, err := f.in.File.CreateFile(ctx, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "File successfully created",
		"Id":      newFile.Id,
		"Url":     newFile.FilePath,
	})
}

func (f File) GetProjectFiles(ctx *gin.Context) {
	projectId := ctx.Query("projectId")

	files, err := f.in.File.GetProjectFiles(ctx, projectId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(files) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no files found for this project",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Files": files,
		})
	}
}

func (f File) CreateSignedURL(ctx *gin.Context) {
	id := ctx.Param("id")

	body := &drivers.File{}

	err := ctx.ShouldBindJSON(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	file := &domain.File{
		Id:       id,
		FileName: body.FileName,
	}

	url, err := f.in.File.CreateSignedURL(ctx, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Id":  id,
		"Url": url,
	})
}
