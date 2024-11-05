package Controllers

import (
	"fmt"
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
	var guid = uuid.New().String()

	// Lấy pageIndex và pageSize từ request query parameters
	pageIndexParam := context.Query("pageIndex")
	pageSizeParam := context.Query("pageSize")

	// Kiểm tra và chuyển đổi pageIndex và pageSize từ string sang int64
	pageIndex, err := strconv.ParseInt(pageIndexParam, 10, 64)
	if err != nil || pageIndex <= 0 {
		pageIndex = 1 // mặc định giá trị pageIndex là 1 nếu không hợp lệ
	}

	pageSize, err := strconv.ParseInt(pageSizeParam, 10, 64)
	if err != nil || pageSize <= 0 {
		pageSize = math.MaxInt //Lấy hết data
	}

	fmt.Println("PageIndex:", pageIndex)
	fmt.Println("PageSize:", pageSize)
	// * Get Kols from the database based on the range of pageIndex and pageSize
	// * TODO: Implement the logic to get parameters from the request
	// ? If parameter passed in the request is not valid, return the response with HTTP Status Bad Request (400)
	// @params: pageIndex
	// @params: pageSize

	// * Perform Logic Here
	// ! Pass the parameters to the Logic Layer
	kols, error := Logic.GetKolLogic(pageIndex, pageSize) //Truyền tham số vào hàm
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
