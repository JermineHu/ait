package handler

import (
	"bufio"
	"bytes"
	"errors"
	. "github.com/JermineHu/ait/consts"
	"github.com/JermineHu/ait/model"
	dm "git.vdo.space/foy/model"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
	. "github.com/JermineHu/ait/controller"
)

type bufferedWriter struct {
	echo.Response
	out    *bufio.Writer
	Buffer bytes.Buffer
}

func (g *bufferedWriter) Write(data []byte) (int, error) {
	g.Buffer.Write(data)
	return g.out.Write(data)
}

func GetRequestHeader(header http.Header, key string) string {
	if values, _ := header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

func AuthorizationHandler(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error{
		//mc := c.MustGet(ControllerFKey).(*FController)
		//token := GetRequestHeader(c.Request.Header, Token)
		//if token == "" {
		//	c.Abort()
		//	panic(dm.NewError(errors.New("Token was needed")))
		//}
		//
		//usr, err := model.GetUserByToken(token)
		//if err != nil || usr == nil {
		//	c.Abort()
		//	panic(dm.NewError(errors.New("User not found")))
		//}
		//
		//mc.User = usr
		//c.Set(ControllerFKey, mc)

		return h(c)
	}
}

func DomainHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet(ControllerFKey).(*FController)
		name := GetRequestHeader(c.Request.Header, "User-Name")
		if name == "" {
			c.Abort()
			panic(dm.NewError(errors.New("User-Name was needed")))
		}
		usr, err := model.GetUserByName(name)
		if err != nil || usr == nil {
			c.Abort()
			panic(dm.NewError(errors.New("User-Name not found")))
		}

		mc.User = usr
		c.Set(ControllerFKey, mc)
	}
}
