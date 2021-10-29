package rest

import (
	"context"
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
		ProjectID: body.ProjectID,
		FileName:  body.FileName,
	}

	newFile, err := f.in.File.CreateFile(ctx, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "File successfully created",
		"id":      newFile.ID,
		"url":     newFile.FilePath,
	})
}

func (f File) GetAllFiles(ctx *gin.Context) {
	files, err := f.in.File.GetFiles(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

func (f File) EditFile(ctx context.Context, id string) error {
	file := &domain.File{
		ID: id,
	}

	err := f.in.File.EditFile(ctx, file)
	if err != nil {
		return err
	}

	return nil
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
		ID:       id,
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
		"ID":  id,
		"url": url,
	})
}
