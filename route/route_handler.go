package route

import (
	"github.com/JermineHu/ait/common"
	. "github.com/JermineHu/ait/consts"
	"github.com/JermineHu/ait/model"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
	"github.com/labstack/echo"
	"github.com/JermineHu/ait/utils"
)

type BodyError struct {
	Success bool                   `json:"success"`
	Error   *model.ErrorMsgDefault `json:"error,omitempty"`
	Data    interface{}            `json:"data,omitempty"`
}

func CheckDBSession (h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c}
		return h(cc)
	}
}

func InitHandler(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//c.Writer.Header().Set("P3P", "CP=CURa ADMa DEVa PSAo PSDo OUR BUS UNI PUR INT DEM STA PRE COM NAV OTC NOI DSP COR")
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, No-Cache, X-Requested-With, If-Modified-Since, Pragma, Last-Modified, Cache-Control, Expires, Content-Type, X-E4M-With,Authentication-Token,User-Name, *")
		//
		//if c.Request.Method == "OPTIONS" {
		//	c.AbortWithStatus(http.StatusOK)
		//}
		//
		//mf := f.NewFController(nil, dc.MG)
		//mb := b.NewBController(nil, dc.MG)
		//c.Set(ControllerFKey, mf)
		//c.Set(ControllerBKey, mb)
		//c.Next()
		return h(c)
	}
}

func HeaderErrorHandler(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context)error {
		c.Set(MeanErrorHandelKey, MeanErrorHandleTypeHeader)
		return h(c)
	}
}

func ErrorHandler(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context)error{

		defer func() {
			if err := recover(); err != nil {
				HandleError(c, err)
			}
		}()
		return h(c)
	}
}

func HandleError(c echo.Context, e interface{}) {
	ht := c.Get(MeanErrorHandelKey).(MeanErrorHandleType)
	log.Print(e)
	debug.PrintStack()
	var err *model.ErrorMsgDefault
	switch er := e.(type) {
	case *model.ErrorMsgDefault:
		err = er
	case *gin.Error:
		if er == nil {
			return
		}
		meta := er.Meta.(gin.H)
		err = model.NewResponseError(meta["message"].(string), meta["code"].(int), 500, utils.ErrorTypeInternal)

	default:
		err = common.ServerErr

	}

	if ht == MeanErrorHandleTypeHeader {
		c.JSON(err.HttpCode, gin.H{
			"code":    err.ErrorCode,
			"message": err.Message,
		})

	} else {
		err := BodyError{
			Success: false,
			Error:   err,
		}

		c.JSON(200, err)
	}
}

func BodyErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(MeanErrorHandelKey, MeanErrorHandelTypeBody)
		c.Next()
	}
}
