package Controllers

import (
	"math"
	"net/http"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetKolsController(context *gin.Context) {
	var KolsVM ViewModels.KolViewModel
	var SearchDTO DTO.GetSearchParam
	var guid = uuid.New().String()

	if err := context.ShouldBindJSON(&SearchDTO); err != nil {
		SearchDTO.PageIndex = 1
		SearchDTO.PageSize = math.MaxInt // If no body => get all
	}

	kols, error := Logic.GetKolLogic(int64(SearchDTO.PageIndex), int64(SearchDTO.PageSize))
	if error != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = error.Error()
		KolsVM.PageIndex = 1
		KolsVM.PageSize = 10
		KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = SearchDTO.PageIndex
	KolsVM.PageSize = SearchDTO.PageSize
	KolsVM.Guid = guid
	KolsVM.KOL = kols
	KolsVM.TotalCount = int64(len(kols))
	context.JSON(http.StatusOK, KolsVM)
}
