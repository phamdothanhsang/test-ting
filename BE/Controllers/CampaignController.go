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

	// * Get Kols from the database based on the range of pageIndex and pageSize
	// * TODO: Implement the logic to get parameters from the request
	// ? If parameter passed in the request is not valid, return the response with HTTP Status Bad Request (400)
	// @params: pageIndex
	// @params: pageSize

	// * Perform Logic Here
	// ! Pass the parameters to the Logic Layer
	q := context.Request.URL.Query()
	pageSize, size_error := strconv.Atoi(q.Get("page_size"))
	pageIndex, index_error := strconv.Atoi(q.Get("page_index"))

	if size_error != nil || index_error != nil {
		var error_mess string = "Invalid page_size/page_index"

		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = error_mess
		KolsVM.PageIndex = pageIndex
		KolsVM.PageSize = pageSize
		KolsVM.Guid = guid
		context.JSON(http.StatusBadRequest, KolsVM)
		return
	}

	kols, error := Logic.GetKolLogic(pageSize, pageIndex)

	if error != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = error.Error()
		KolsVM.PageIndex = pageIndex // * change this to the actual page index from the request
		KolsVM.PageSize = pageSize   // * change this to the actual page size from the request
		KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	// * Return the response after the logic is executed
	// ? If the logic is successful, return the response with HTTP Status OK (200)
	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = pageIndex // * change this to the actual page index from the request
	KolsVM.PageSize = pageSize   // * change this to the actual page size from the request
	KolsVM.Guid = guid
	KolsVM.KOL = kols
	KolsVM.TotalCount = int64(len(kols))
	context.JSON(http.StatusOK, KolsVM)
}
