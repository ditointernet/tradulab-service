package rest

import (
	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/gin-gonic/gin"
)

type File struct {
}

func MustNewFile() File {
	return File{}
}

func (f File) CreateFile(ctx *gin.Context) {

	var checker services.File = services.Path{P: "file.csv"}
	_, err := checker.CheckFile()

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"message": "Upload complete",
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
