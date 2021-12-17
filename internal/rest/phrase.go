package rest

import (
	"net/http"
	"strconv"

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
		"Id":      phrase.Id,
		"FileId":  phrase.FileId,
		"Key":     phrase.Key,
		"Content": phrase.Content,
	})
}

func (p Phrase) GetFilePhrases(ctx *gin.Context) {
	fileId := ctx.Query("fileId")
	page := ctx.Query("page")

	numberPage, err := strconv.Atoi(page)
	if err != nil {
		if page == "" {
			numberPage = 1
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Page number must be integer numeric value",
			})
			return
		}
	}

	if numberPage <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Must be bigger than zero",
		})
		return
	}

	phrases, total, err := p.pService.GetFilePhrases(ctx, fileId, numberPage)
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
	ctx.Header("x-total-count", strconv.Itoa(total))
	ctx.JSON(http.StatusOK, gin.H{
		"Phrases": phrases,
	})
}
