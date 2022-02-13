package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *api) GetAllDoctors(ctx *gin.Context) {
	doctors, err := a.storage.Doctor().GetAll()
	if err != nil {
		errorAbort(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if doctors == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, doctors)
}

func (a *api) GetDoctorById(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	doctor, err := a.storage.Doctor().FindById(id)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if doctor == nil {
		emptyObject(ctx)
		return
	}

	ctx.JSON(http.StatusOK, doctor)
}

func (a *api) GetAllDoctorsBySpecialty(ctx *gin.Context) {
	specialty, err := getStringParam(ctx, "specialty")
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	doctors, err := a.storage.Doctor().FindAllBySpecialty(specialty)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if doctors == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, doctors)
}

func (a *api) DeleteDoctorById(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.storage.Doctor().Remove(id); err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
