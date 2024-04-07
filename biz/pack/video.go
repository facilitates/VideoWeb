package pack

import (
	"strconv"
	"videoweb/biz/model/videoweb"
	"videoweb/dao"
)

func BuildVideoFeedResp(videos []dao.VideoDB) videoweb.VideoFeedResponse {
	var videosResp []*videoweb.Video
	for _, video := range videos {
		buildVideo := BuildVideo(video)
		videosResp = append(videosResp, buildVideo)
	}
	return videoweb.VideoFeedResponse{
		Base: &videoweb.BaseResp{
			Code: 200,
			Msg:  "获取视频成功",
		},
		Items: videosResp,
	}
}

func BuildVideoListResp(videos []dao.VideoDB, count int) videoweb.VideoListResponse {
	var videosResp []*videoweb.Video
	for _, video := range videos {
		buildVideo := BuildVideo(video)
		videosResp = append(videosResp, buildVideo)
	}
	return videoweb.VideoListResponse{
		Base: &videoweb.BaseResp{
			Code: 200,
			Msg:  "获取视频成功",
		},
		Data:  videosResp,
		Total: int64(count),
	}
}

func BuildVideo(video dao.VideoDB) *videoweb.Video {
	id := strconv.FormatUint(uint64(video.ID), 10)
	userid := strconv.FormatUint(uint64(video.UserId), 10)
	likecount := strconv.FormatInt(int64(video.LikeCount), 10)
	return &videoweb.Video{
		Id:           id,
		UserId:       userid,
		VideoUrl:     video.VideoUrl,
		CoverUrl:     video.CoverUrl,
		Title:        video.Title,
		Description:  video.Description,
		VisitCount:   "",
		LikeCount:    likecount,
		CommentCount: "",
		CreatedAt:    video.CreatedAt.String(),
		UpdatedAt:    video.UpdatedAt.String(),
	}
}

func BuildPopularVideoResp(videos []dao.VideoDB) videoweb.PopularVideoResponse {
	var videosResp []*videoweb.Video
	for _, video := range videos {
		buildVideo := BuildVideo(video)
		videosResp = append(videosResp, buildVideo)
	}
	return videoweb.PopularVideoResponse{
		Base: &videoweb.BaseResp{
			Code: 200,
			Msg:  "获取视频成功",
		},
		Data: videosResp,
	}
}
