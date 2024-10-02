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

	// * Get Kols from the database based on the range of pageIndex and pageSize
	// * TODO: Implement the logic to get parameters from the request
	// ? If parameter passed in the request is not valid, return the response with HTTP Status Bad Request (400)
	// @params: pageIndex
	// @params: pageSize

	// * Perform Logic Here
	// ! Pass the parameters to the Logic Layer

	if err := context.ShouldBindJSON(&SearchDTO); err != nil {
		SearchDTO.PageIndex = 1
		SearchDTO.PageSize = math.MaxInt // If no body => get all
	}

	kols, error := Logic.GetKolLogic(int64(SearchDTO.PageIndex), int64(SearchDTO.PageSize))
	if error != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = error.Error()
		KolsVM.PageIndex = 1 // * change this to the actual page index from the request
		KolsVM.PageSize = 10 // * change this to the actual page size from the request
		KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	// * Return the response after the logic is executed
	// ? If the logic is successful, return the response with HTTP Status OK (200)
	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = SearchDTO.PageIndex // * change this to the actual page index from the request
	KolsVM.PageSize = SearchDTO.PageSize   // * change this to the actual page size from the request
	KolsVM.Guid = guid
	KolsVM.KOL = kols
	KolsVM.TotalCount = int64(len(kols))
	context.JSON(http.StatusOK, KolsVM)
}
