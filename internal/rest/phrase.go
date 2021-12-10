package rest

import (
	"net/http"

	"github.com/ditointernet/tradulab-service/internal/core/services"
	"github.com/gin-gonic/gin"
)

type Phrase struct {
	pService *services.Phrase
}

func MustNewPhrase(pService *services.Phrase) Phrase {
	return Phrase{
		pService: pService,
	}
}

func (p Phrase) GetPhrasesById(ctx *gin.Context) {
	phraseId := ctx.Param("id")

	phrase, err := p.pService.GetPhrasesById(ctx, phraseId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":      phrase.ID,
		"fileId":  phrase.FileID,
		"key":     phrase.Key,
		"content": phrase.Content,
	})
}

func (p Phrase) GetFilePhrases(ctx *gin.Context) {
	fileId := ctx.Query("fileId")
	page := ctx.Query("page")

	phrases, err := p.pService.GetFilePhrases(ctx, fileId, page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(phrases) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no phrases found for this file in this page",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"phrases": phrases,
	})
}
