package pack

import (
	"strconv"
	"videoweb/biz/model/videoweb"
	"videoweb/dao"
)

func BuildLikeList(videolist []dao.VideoDB) videoweb.LikeListResponse {
	var videosResp []*videoweb.Video
	i := 0
	for _, video := range videolist {
		buildVideo := BuildVideo(video)
		videosResp = append(videosResp, buildVideo)
		i++
	}
	if i != 0 {
		return videoweb.LikeListResponse{
			Base: &videoweb.BaseResp{
				Code: 200,
				Msg:  "获取视频成功",
			},
			Data: videosResp,
		}
	} else {
		return videoweb.LikeListResponse{
			Base: &videoweb.BaseResp{
				Code: 200,
				Msg:  "用户没有点赞任何视频",
			},
		}
	}
}

func BuildCommentListResp(commentlist []dao.CommentDB) videoweb.CommentListResponse {
	var commentList []*videoweb.Comment
	for _, comment := range commentlist {
		var buildcomment *videoweb.Comment
		buildcomment = BuildComment(comment)
		commentList = append(commentList, buildcomment)
	}
	return videoweb.CommentListResponse{
		Base: &videoweb.BaseResp{
			Code: 200,
			Msg:  "获取评论列表成功",
		},
		Data: commentList,
	}
}

func BuildComment(comment dao.CommentDB) *videoweb.Comment {
	commentid := strconv.FormatUint(uint64(comment.ID), 10)
	userid := strconv.FormatUint(uint64(comment.UserId), 10)
	videoid := strconv.FormatUint(uint64(comment.VideoId), 10)
	parentid := strconv.FormatUint(uint64(comment.ParentId), 10)
	return &videoweb.Comment{
		ID:        commentid,
		UserID:    userid,
		VideoID:   videoid,
		ParentID:  parentid,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.String(),
		UpdatedAt: comment.UpdatedAt.String(),
	}
}

func BuildLikeActionResp(state int) videoweb.LikeActionResponse {
	if state == 1 {
		return videoweb.LikeActionResponse{
			Base: &videoweb.BaseResp{
				Code: 200,
				Msg:  "点赞成功",
			},
		}
	} else {
		return videoweb.LikeActionResponse{
			Base: &videoweb.BaseResp{
				Code: 200,
				Msg:  "取消点赞成功",
			},
		}
	}
}
