package ocr

import (
	"ITG/internal/common"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OCRHandler struct {
	svc *OCRService
}

func NewOCRHandler(svc *OCRService) *OCRHandler {
	return &OCRHandler{
		svc: svc,
	}
}
func (h *OCRHandler) HandleUpload(do func(text []string) (any, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		reader, err := c.Request().MultipartReader()
		if err != nil {
			return common.Error(c, err)
		}
		texts, err := h.svc.ProcessImages(c, reader)
		if err != nil {
			return common.Error(c, err)
		}
		result, err := do(texts)
		if err != nil {
			return common.Error(c, err)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "success",
			"results": result,
		})
	}
}
