// This file "echo.go" is created by Lincan Li at 5/9/16.
// Copyright © 2016 - Lincan Li. All rights reserved

package services

import (
	"git.ngs.tech/mean/icarus/model"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

type Icarus struct {
	MDB *mgo.Database
}

const MDB_CONTEXT = "MDB_CTX"
const C_CONTEXT = "C_CONTEXT"

func (a *Icarus) Context(ctx context.Context) {
	if ctx == nil {
		return
	}
	a.MDB = ctx.Value(MDB_CONTEXT).(*mgo.Database)
	model.InitMDB(a.MDB)
}
