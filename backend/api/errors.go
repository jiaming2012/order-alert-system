package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jiaming2012/order-alert-system/backend/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

func sendBadServerErrResponse(err error, ctx *gin.Context) {
	logrus.Error(err)
	ctx.Status(http.StatusInternalServerError)
}

func sendBadServerHtmlResponse(err error, ctx *gin.Context) {
	logrus.Error(err)
	ctx.HTML(http.StatusInternalServerError, "500-error.html", gin.H{})
}

func sendBadRequestHtmlResponse(err error, ctx *gin.Context) {
	encodedErrorMsg := url.QueryEscape(err.Error())
	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/400-error.html?errorMsg=%s", encodedErrorMsg))
}

func sendBadRequestErrResponse(errType string, err error, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, models.BadResponseError{
		Type: errType,
		Msg:  err.Error(),
	})
}
