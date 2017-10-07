package services

import (
	dm "git.ngs.tech/mean/daniel/model"
	. "git.ngs.tech/mean/daniel/proto"
	"git.ngs.tech/mean/daniel/utils"
	"git.ngs.tech/mean/icarus/model"
	"github.com/micro/protobuf/proto"
	"golang.org/x/net/context"
)

//NewMoods 批量创建Mood
func (a Icarus) NewMoods(ctx context.Context, req *MoodList, rsp *MoodList) error {
	a.Context(ctx)
	ml, err := model.NewMoods(dm.Echos2Moods(req.Data))

	rsp.Data = dm.Moods2Echos(ml)

	return err

}

//DeleteMoodsByIDs 根据数据的id数组删除数据
func (a Icarus) DeleteMoodsByIDs(ctx context.Context, req *StringIDs, rsp *Bool) error {
	a.Context(ctx)

	rsp.Bool = model.DeleteMoodsByIDs(utils.Strs2ObjectIds(req.Data))
	return nil
}

//GetMoodByIDs 根据数据的id数组查询数据
func (a Icarus) GetMoodByIDs(ctx context.Context, req *StringIDs, rsp *MoodList) error {
	a.Context(ctx)

	ml, err := model.GetMoodByIDs(utils.Strs2ObjectIds(req.Data))

	rsp.Data = dm.Moods2Echos(ml)

	return err

}

//GetMoodByNum 根据数据的num查询数据
func (a Icarus) GetMoodByNum(ctx context.Context, req *Mood, rsp *MoodList) error {
	a.Context(ctx)

	ml, err := model.GetMoodByNum(req.Number)

	rsp.Data = dm.Moods2Echos(ml)

	return err

}

//UpdateMoods 批量更新Mood
func (a Icarus) UpdateMoods(ctx context.Context, req *MoodList, rsp *MoodList) error {
	a.Context(ctx)

	ml, err := model.UpdateMoodsByNum(dm.Echos2Moods(req.Data))

	rsp.Data = dm.Moods2Echos(ml)

	return err

}

//FindMoodsByParams 根据查询参数获取Mood的分页数据
func (a Icarus) FindMoodsByParams(ctx context.Context, req *RequestParams, rsp *PageModel) error {
	a.Context(ctx)

	news, count, err := model.FindMoodsByParams(req)

	if err != nil {
		return err
	}

	ns := [][]byte{}
	for _, x := range news {

		u, err := proto.Marshal(dm.Mood2Echo(x))

		if err != nil {

			return err
		}

		ns = append(ns, u)
	}
	rsp.TotalCount = count
	rsp.List = ns

	return err

}
