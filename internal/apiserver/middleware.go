package apiserver

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Params.Get("id")

	if !ok {
		return -1, errors.New("id was not found")
	}

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return -1, errors.New("id is not numerical")
	}

	return idInt, nil
}

func getStringParam(ctx *gin.Context, param string) (string, error) {
	value, ok := ctx.Params.Get(param)

	if !ok {
		return "", errors.New(param + " was not found")
	}

	return value, nil
}
