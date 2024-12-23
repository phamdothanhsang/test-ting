package Controllers

import (
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
	var guid = uuid.New().String()

	// Get pageIndex and pageSize from query parameters
	pageIndex, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil || pageIndex < 1 {
		pageIndex = 1
	}

	pageSize, err := strconv.Atoi(context.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Fetch paginated KOLs and total count
	kols, totalCount, fetchErr := Logic.GetKolLogic(pageIndex-1, pageSize)
	if fetchErr != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = fetchErr.Error()
		KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	// Calculate totalPages from totalCount and pageSize
	totalPages := (totalCount + int64(pageSize) - 1) / int64(pageSize)

	// Return paginated data
	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.TotalCount = totalCount
	KolsVM.TotalPages = totalPages
	KolsVM.KolInformation = kols
	KolsVM.Guid = guid

	context.JSON(http.StatusOK, KolsVM)
}
