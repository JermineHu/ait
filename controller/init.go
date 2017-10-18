package controller

import (
	"gopkg.in/mgo.v2"
	"github.com/JermineHu/ait/common"
)

type FController struct {
	common.Controller
}

func NewFController(MDBSession *mgo.Session, MDB *mgo.Database) *FController {
	mc := &FController{}
	//mc.MDB = MDB
	return mc
}
