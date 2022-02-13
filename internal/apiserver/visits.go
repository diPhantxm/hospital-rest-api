package apiserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *api) GetAllVisists(ctx *gin.Context) {
	visits, err := a.storage.Visit().GetAll()
	if err != nil {
		errorAbort(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if visits == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, visits)
}

func (a *api) GetVisitById(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	visit, err := a.storage.Visit().FindById(id)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if visit == nil {
		emptyObject(ctx)
		return
	}

	ctx.JSON(http.StatusOK, visit)
}

func (a *api) GetAllVisitsByPatientId(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	visits, err := a.storage.Visit().FindAllByPatientId(id)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if visits == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, visits)
}

func (a *api) GetAllVisitsByDiseaseId(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	visits, err := a.storage.Visit().FindAllByDiseaseId(id)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if visits == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, visits)
}

func (a *api) GetAllVisitsByDoctorId(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	visits, err := a.storage.Visit().FindAllByDoctorId(id)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if visits == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, visits)
}

func (a *api) GetAllVisitsByDate(ctx *gin.Context) {
	date, err := getStringParam(ctx, "date")
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	dateParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	visits, err := a.storage.Visit().FindAllByDate(dateParsed)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if visits == nil {
		emptySlice(ctx)
		return
	}

	ctx.JSON(http.StatusOK, visits)
}

func (a *api) DeleteVisitById(ctx *gin.Context) {
	id, err := getId(ctx)
	if err != nil {
		errorAbort(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.storage.Visit().Remove(id); err != nil {
		errorAbort(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
