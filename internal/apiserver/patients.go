package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *api) GetAllPatients(ctx *gin.Context) {
	patients, err := a.storage.Patient().GetAll()
	if err != nil {
		errorAbort(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if patients == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, patients)
}

func (a *api) GetPatientById(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
	}

	patient, err := a.storage.Patient().FindById(id)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
	}
	if patient == nil {
		emptyObject(ctx)
		return
	}

	ctx.JSON(http.StatusOK, patient)
}

func (a *api) DeletePatientById(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
	}

	if err := a.storage.Patient().Remove(id); err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
