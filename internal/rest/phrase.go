package rest

import (
	"net/http"

	"github.com/ditointernet/tradulab-service/drivers"
	"github.com/gin-gonic/gin"
)

type Phrase struct {
}

func MustNewPhrase() Phrase {
	return Phrase{}
}

func (p Phrase) FindByID(ctx *gin.Context) {
	ID, _ := ctx.Params.Get("id")

	ctx.JSON(http.StatusOK, drivers.Phrase{ID: ID, Key: "welcome.hello"})
}
