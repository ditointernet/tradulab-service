package rest

import (
	"context"
	"net/http"

	"github.com/ditointernet/tradulab-service/internal/core/domain"
	"github.com/gin-gonic/gin"
)

type PhraseService interface {
	GetByID(context.Context, string) (domain.Phrase, error)
}

type Phrase struct {
	Service PhraseService
}

func MustNewPhrase(service PhraseService) Phrase {
	return Phrase{
		Service: service,
	}
}

func (p Phrase) FindByID(ctx *gin.Context) {
	ID, _ := ctx.Params.Get("id")

	phrase, err := p.Service.GetByID(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, phrase)
}
