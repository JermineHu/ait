package services

import (
	dm "git.ngs.tech/mean/daniel/model"
	. "git.ngs.tech/mean/daniel/proto"
	"git.ngs.tech/mean/daniel/utils"
	"git.ngs.tech/mean/icarus/model"
	"github.com/micro/protobuf/proto"
	"golang.org/x/net/context"
)

//NewNatures 批量创建Nature
func (a Icarus) NewNatures(ctx context.Context, req *NatureList, rsp *NatureList) error {
	a.Context(ctx)

	ml, err := model.NewNatures(dm.Echos2Natures(req.Data))

	rsp.Data = dm.Natures2Echos(ml)

	return err

}

//DeleteNaturesByIDs 根据数据的id数组删除数据
func (a Icarus) DeleteNaturesByIDs(ctx context.Context, req *StringIDs, rsp *Bool) error {
	a.Context(ctx)

	rsp.Bool = model.DeleteNaturesByIDs(utils.Strs2ObjectIds(req.Data))
	return nil
}

//GetNatureByIDs 根据数据的id数组查询数据
func (a Icarus) GetNatureByIDs(ctx context.Context, req *StringIDs, rsp *NatureList) error {
	a.Context(ctx)

	ml, err := model.GetNatureByIDs(utils.Strs2ObjectIds(req.Data))

	rsp.Data = dm.Natures2Echos(ml)

	return err

}

//UpdateNatures 批量更新Nature
func (a Icarus) UpdateNatures(ctx context.Context, req *NatureList, rsp *NatureList) error {
	a.Context(ctx)

	ml, err := model.UpdateNaturesByIDs(dm.Echos2Natures(req.Data))

	rsp.Data = dm.Natures2Echos(ml)

	return err

}

//FindNaturesByParams 根据查询参数获取Nature的分页数据
func (a Icarus) FindNaturesByParams(ctx context.Context, req *RequestParams, rsp *PageModel) error {
	a.Context(ctx)

	list, count, err := model.FindNaturesByParams(req)

	if err != nil {
		return err
	}

	ns := [][]byte{}
	for _, x := range list {

		u, err := proto.Marshal(dm.Nature2Echo(x))

		if err != nil {

			return err
		}

		ns = append(ns, u)
	}
	rsp.TotalCount = count
	rsp.List = ns

	return err

}
