package common

import (
	ml "github.com/JermineHu/ait/model"
	"git.vdo.space/foy/model"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"time"
)

type Controller struct {
	RDB    *gorm.DB
	MDB    *mgo.Database
	User   *ml.User
	Broker broker.Broker
	//Logger *log.MeanLogger
}

func NewMeanController(MDBSession *mgo.Session, MDB *mgo.Database) *Controller {
	mc := &Controller{
		//RDB: RDB,
		MDB: MDB,
	}
	return mc
}

type SingleEntity struct {
	Validation model.NullBool `json:"validation,omitempty"`
	QNToken    string         `json:"QN-Token,omitempty"`
	Expires    *time.Time     `json:"QN-Expires,omitempty"`
}

const (
	CONSUL_ADDRESSES_KEY = "CONSUL_ADDRS"
)

const (
	DREAM_SERVICE_KEY  = "tech-ngs-dream"
	ATHENA_SERVICE_KEY = "tech-ngs-athena"
)


type TimeWrapper struct {
	client.Client
}

func (l *TimeWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	start := time.Now()

	defer func() {
		end := time.Now()
		end.Sub(start)
	}()

	return l.Client.Call(ctx, req, rsp)
}

func Wrapper(c client.Client) client.Client {
	return &TimeWrapper{c}
}