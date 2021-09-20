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

func MustNewFile(in ServiceInput) (*File, error) {
	if in.File == nil {
		return nil, fmt.Errorf("Error message...")
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

	err = f.in.File.SaveFile(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = f.in.File.CheckFile(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	file := domain.File{
		ProjectID: body.ProjectID,
		FilePath:  body.FilePath,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Upload complete",
		"file":    file,
	})

	/*form, _ := ctx.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		body := &File{}
		err := ctx.ShouldBindJSON(body)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.SaveUploadedFile(file, "/ports")
		body.FilePath = "/ports/"
	}

	jsonValue, err := json.Marshal(body)*/
}
