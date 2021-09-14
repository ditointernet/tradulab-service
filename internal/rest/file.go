package rest

import (
	"net/http"

	"github.com/ditointernet/tradulab-service/drivers"
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/gin-gonic/gin"
)

type File struct {
}

func MustNewFile() File {
	return File{}
}

func (f File) CreateFile(ctx *gin.Context) {

	var i services.File = services.Path{"file.csv"}
	i.CheckFile()

	ctx.JSON(http.StatusOK, drivers.File{ID: i.ID, ProjectID: i.ProjectID, FilePath: i.FilePath})

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
