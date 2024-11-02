package Controllers

import (
	"net/http"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetKolsController(context *gin.Context) {
	var KolsVM ViewModels.KolViewModel
	var guid = uuid.New().String()

	// * Get Kols from the database based on the range of pageIndex and pageSize
	// * TODO: Implement the logic to get parameters from the request
	// ? If parameter passed in the request is not valid, return the response with HTTP Status Bad Request (400)
	// @params: pageIndex
	// @params: pageSize
	type PageParams struct {
		PageIndex int `form:"pageIndex" binding:"required"`
		PageSize int `form:"pageSize" binding:"required"`
	}
	var pageParams PageParams
	if context.ShouldBindQuery(&pageParams) != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	// * Perform Logic Here
	// ! Pass the parameters to the Logic Layer
	kols, error := Logic.GetKolLogic(pageParams.PageIndex, pageParams.PageSize)
	if error != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = error.Error()
		KolsVM.PageIndex = int64(pageParams.PageIndex) // * change this to the actual page index from the request
		KolsVM.PageSize = int64(pageParams.PageSize) // * change this to the actual page size from the request
		KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	// * Return the response after the logic is executed
	// ? If the logic is successful, return the response with HTTP Status OK (200)
	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = int64(pageParams.PageIndex) // * change this to the actual page index from the request
	KolsVM.PageSize = int64(pageParams.PageSize) // * change this to the actual page size from the request
	KolsVM.Guid = guid
	KolsVM.KOL = kols
	KolsVM.TotalCount = int64(len(kols))
	context.JSON(http.StatusOK, KolsVM)
}
