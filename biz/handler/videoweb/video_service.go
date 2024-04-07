// Code generated by hertz generator.

package videoweb

import (
	"context"
	"errors"
	"mime/multipart"
	"strconv"
	"time"
	"videoweb/biz/pack"
	"videoweb/dao"
	"videoweb/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"videoweb/biz/model/videoweb"
)

// VideoFeed .
// @router /video/feed [GET]
func VideoFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req videoweb.VideoFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, req.LatestTime)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var video []dao.VideoDB
	video, err = req.VideoFeed(parsedTime)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
	} else {
		resp := pack.BuildVideoFeedResp(video)
		c.JSON(consts.StatusOK, resp)
	}
}

// UploadVideo .
// @router /video/publish [POST]
func UploadVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req videoweb.UploadVideoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var file *multipart.FileHeader
	file, err = c.FormFile("data")
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
	}
	if utils.ParseVideoExt(file.Filename) {
		c.String(consts.StatusBadRequest, errors.New("文件格式错误").Error())
	}
	claim, _ := utils.ParseToken(string(c.GetHeader("Access-Token")))
	filepath := "./upload/video/" + claim.UserName + "/" + file.Filename
	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
	}
	err = req.UploadVideo(claim, filepath)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
	}
	resp := &videoweb.UploadVideoResponse{
		Base: &videoweb.BaseResp{
			Code: 200,
			Msg:  "视频上传成功",
		},
	}
	c.JSON(consts.StatusOK, resp)
}

// VideoList .
// @router /video/list [GET]
func VideoList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req videoweb.VideoListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var videolist []dao.VideoDB
	userid, _ := strconv.ParseUint(req.UserID, 10, 64)
	var count int
	videolist, count, err = req.VideoList(uint(userid))
	var videoList []dao.VideoDB
	start := int(req.PageNum * req.PageSize)
	for i, video := range videolist {
		if i >= start && i < start+int(req.PageSize) {
			videoList = append(videoList, video)
		}
	}
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
	} else {
		resp := pack.BuildVideoListResp(videoList, count)
		c.JSON(consts.StatusOK, resp)
	}
}

// PopularVideo .
// @router /video/popular [GET]
func PopularVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req videoweb.PopularVideoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	videos, err := req.PopularVideo()
	if err != nil {
		resp := pack.BuildErrResp(400, err)
		c.JSON(200, resp)
	}
	resp := pack.BuildPopularVideoResp(videos)
	c.JSON(consts.StatusOK, resp)
}

// SearchVideo .
// @router /video/search [POST]
func SearchVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req videoweb.SearchVideoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(videoweb.SearchVideoResponse)

	c.JSON(consts.StatusOK, resp)
}
