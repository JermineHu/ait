package services

import (
	dm "git.ngs.tech/mean/daniel/model"
	. "git.ngs.tech/mean/daniel/proto"
	"git.ngs.tech/mean/daniel/utils"
	"git.ngs.tech/mean/icarus/model"
	"github.com/micro/protobuf/proto"
	"golang.org/x/net/context"
)

//NewImageFilters 批量创建ImageFilter
func (a Icarus) NewImageFilters(ctx context.Context, req *ImageFilterList, rsp *ImageFilterList) error {
	a.Context(ctx)

	ml, err := model.NewImageFilters(dm.Echos2ImageFilters(req.Data))

	rsp.Data = dm.ImageFilters2Echos(ml)

	return err

}

//DeleteImageFiltersByIDs 根据数据的id数组删除数据
func (a Icarus) DeleteImageFiltersByIDs(ctx context.Context, req *StringIDs, rsp *Bool) error {
	a.Context(ctx)

	rsp.Bool = model.DeleteImageFiltersByIDs(utils.Strs2ObjectIds(req.Data))
	return nil
}

//GetImageFilterByIDs 根据数据的id数组查询数据
func (a Icarus) GetImageFilterByIDs(ctx context.Context, req *StringIDs, rsp *ImageFilterList) error {
	a.Context(ctx)

	ml, err := model.GetImageFilterByIDs(utils.Strs2ObjectIds(req.Data))

	rsp.Data = dm.ImageFilters2Echos(ml)

	return err

}

//GetImageFilterByNum 根据数据的num查询数据
func (a Icarus) GetImageFilterByNum(ctx context.Context, req *ImageFilter, rsp *ImageFilterList) error {
	a.Context(ctx)

	ml, err := model.GetImageFilterByNum(req.Number)

	rsp.Data = dm.ImageFilters2Echos(ml)

	return err

}

//UpdateImageFilters 批量更新ImageFilter
func (a Icarus) UpdateImageFilters(ctx context.Context, req *ImageFilterList, rsp *ImageFilterList) error {
	a.Context(ctx)

	ml, err := model.UpdateImageFilters(dm.Echos2ImageFilters(req.Data))

	rsp.Data = dm.ImageFilters2Echos(ml)

	return err

}

//FindImageFiltersByParams 根据查询参数获取ImageFilter的分页数据
func (a Icarus) FindImageFiltersByParams(ctx context.Context, req *RequestParams, rsp *PageModel) error {
	a.Context(ctx)

	news, count, err := model.FindImageFiltersByParams(req)

	if err != nil {
		return err
	}

	ns := [][]byte{}
	for _, x := range news {

		u, err := proto.Marshal(dm.ImageFilter2Echo(x))

		if err != nil {

			return err
		}

		ns = append(ns, u)
	}
	rsp.TotalCount = count
	rsp.List = ns

	return err

}
