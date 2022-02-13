package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *api) GetAllDiseases(ctx *gin.Context) {
	diseases, err := a.storage.Disease().GetAll()

	if err != nil {
		errorAbort(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if diseases == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, diseases)
}

func (a *api) GetDiseaseById(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	disease, err := a.storage.Disease().FindById(id)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if disease == nil {
		emptyObject(ctx)
		return
	}

	ctx.JSON(http.StatusOK, disease)
}

func (a *api) GetAllDiseasesByPatientId(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	diseases, err := a.storage.Disease().FindAllByPatientId(id)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if diseases == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, diseases)
}

func (a *api) GetAllDiseasesByName(ctx *gin.Context) {
	name, err := getStringParam(ctx, "name")
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	diseases, err := a.storage.Disease().FindAllByName(name)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if diseases == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, diseases)
}

func (a *api) DeleteDiseaseById(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.storage.Disease().Remove(id); err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
