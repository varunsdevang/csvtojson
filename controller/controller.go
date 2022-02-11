package controller

import (
	"github/csvtojson/converter"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	//
	ConverterController struct {
		Converter converter.Converter
	}

	ErrorResponse struct {
		Message string `json:"message"`
	}
)

func NewConverterController() *ConverterController {
	return &ConverterController{Converter: &converter.CsvToJsonConverter{}}
}

// CsvToJsonHandler
// @Summary      handler for csv to json conversion
// @Description  generates json array for provided csv file
// @Accept       mpfd
// @Produce      json
// @Success      200  {object}  []map[string]string
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /csvtojson [post]
func (c *ConverterController) CsvToJsonHandler(cxt echo.Context) error {
	err := cxt.Request().ParseMultipartForm(1024 * 1024)
	if err != nil {
		return cxt.JSON(http.StatusBadRequest, ErrorResponse{Message: "request body should be multipart/form-data"})
	}

	if len(cxt.Request().MultipartForm.File["csv"]) == 0 {
		return cxt.JSON(http.StatusBadRequest, ErrorResponse{Message: "request does not contain csv file"})
	}

	f, err := cxt.Request().MultipartForm.File["csv"][0].Open()
	if err != nil {
		return cxt.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid file"})
	}

	defer f.Close()

	output, errMsg, err := c.Converter.CsvToJson(f)
	if err != nil {
		log.Error("error : ", err)
		log.Error("message : ", errMsg)
		return cxt.JSON(c.errToHTTPCode(err), ErrorResponse{Message: errMsg})
	}

	return cxt.JSON(http.StatusOK, output)
}

func (c *ConverterController) errToHTTPCode(err error) int {
	switch err {
	case converter.ErrorInvalidHeader,
		converter.ErrorInvalidRow:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
