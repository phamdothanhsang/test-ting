package Controllers

import (
	"math"
	"net/http"
	"strconv"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetKolsController(context *gin.Context) {
	var KolsVM ViewModels.KolViewModel
	guid := uuid.New().String()

	pageIndex := 1
	pageSize := math.MaxInt // Default to max int for all records

	if pageIndexParam := context.Query("pageIndex"); pageIndexParam != "" {
		if idx, err := strconv.Atoi(pageIndexParam); err == nil {
			pageIndex = idx
		}
	}

	if pageSizeParam := context.Query("pageSize"); pageSizeParam != "" {
		if size, err := strconv.Atoi(pageSizeParam); err == nil {
			pageSize = size
		}
	}

	// Fetch KOL data using logic layer
	kols, err := Logic.GetKolLogic(int64(pageIndex), int64(pageSize))
	if err != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = err.Error()
		KolsVM.PageIndex = 1
		KolsVM.PageSize = 10
		KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = int64(pageIndex)
	KolsVM.PageSize = int64(pageSize)
	KolsVM.Guid = guid
	KolsVM.KOL = kols
	KolsVM.TotalCount = int64(len(kols))
	context.JSON(http.StatusOK, KolsVM)
}
