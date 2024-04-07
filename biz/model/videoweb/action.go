package videoweb

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
	"videoweb/dao"
	"videoweb/pkg/utils"
)

func (req *LikeActionRequest) LikeAction(userid string) (int, error) {
	if req.ActionType == "" {
		return 0, errors.New("点赞类型错误")
	} else if req.ActionType == "1" {
		if req.CommentID == "" && req.VideoID == "" {
			return 0, errors.New("未指定点赞目标")
		} else if req.CommentID != "" && req.VideoID != "" {
			return 0, errors.New("只能指定一个点赞目标")
		} else if req.CommentID != "" && req.VideoID == "" {
			var commentid string
			commentid = "comment:" + req.VideoID
			isLiked, err := dao.Redisdb.SIsMember(commentid, userid).Result()
			if err != nil {
				return 0, err
			}
			if isLiked {
				return 0, errors.New("已经点过赞了")
			} else {
				err := dao.Redisdb.SAdd(commentid, userid).Err()
				if err != nil {
					return 0, err
				}
				return 1, nil
			}
		} else if req.CommentID == "" && req.VideoID != "" {
			var videoid string
			videoid = "video:" + req.VideoID
			isLiked, err := dao.Redisdb.SIsMember(videoid, userid).Result()
			if err != nil {
				return 0, err
			}
			if isLiked {
				return 0, errors.New("已经点过赞了")
			} else {
				err := dao.Redisdb.ZIncrBy("videolike", 1, userid).Err()
				if err != nil {
					return 0, err
				}
				err = dao.Redisdb.SAdd(videoid, userid).Err()
				if err != nil {
					return 0, err
				}
				userid = "user:" + userid
				err = dao.Redisdb.SAdd(userid, req.VideoID).Err()
				if err != nil {
					return 0, err
				}
				return 1, nil
			}
		}
		return 1, nil
	} else if req.ActionType == "2" {
		if req.CommentID == "" && req.VideoID == "" {
			return 0, errors.New("未指定取消点赞目标")
		} else if req.CommentID != "" && req.VideoID != "" {
			return 0, errors.New("只能指定一个取消点赞目标")
		} else if req.CommentID != "" && req.VideoID == "" {
			var commentid string
			commentid = "comment:" + req.VideoID
			isLiked, err := dao.Redisdb.SIsMember(commentid, userid).Result()
			if err != nil {
				return 0, err
			}
			if isLiked {
				err := dao.Redisdb.SRem(commentid, userid).Err()
				if err != nil {
					return 0, err
				}
				return 2, nil
			} else {
				return 0, errors.New("未点赞")
			}
		} else if req.CommentID == "" && req.VideoID != "" {
			var videoid string
			videoid = "video:" + req.VideoID
			isLiked, err := dao.Redisdb.SIsMember(videoid, userid).Result()
			if err != nil {
				return 0, err
			}
			if isLiked {
				err := dao.Redisdb.ZIncrBy("videolike", 1, userid).Err()
				if err != nil {
					return 0, err
				}
				err = dao.Redisdb.SRem(videoid, userid).Err()
				if err != nil {
					return 0, err
				}
				userid = "user:" + userid
				err = dao.Redisdb.SRem(userid, req.VideoID).Err()
				if err != nil {
					return 0, err
				}
				return 2, err
			} else {
				return 0, errors.New("未点赞")
			}
		}
		return 2, nil
	}
	return 2, nil
}

func (req *LikeListRequest) LikeList(userid string, useriddb uint) ([]dao.VideoDB, error) {
	var videoList []dao.VideoDB
	err := dao.DB.Model(&dao.UserDB{}).Where("id=?", useriddb).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return videoList, errors.New("该用户不存在")
		} else if !gorm.IsRecordNotFoundError(err) {
			return videoList, err
		}
	}
	videos, err := dao.Redisdb.SMembers(userid).Result()
	if err != nil {
		return videoList, err
	}
	start := int(req.PageNum * req.PageSize)
	for i, video := range videos {
		if i >= start && i <= start+int(req.PageSize) {
			var videoDB dao.VideoDB
			videoid, _ := strconv.ParseUint(video, 10, 64)
			err = dao.DB.First(&videoDB, uint(videoid)).Error
			if err != nil {
				return videoList, err
			}
			videoList = append(videoList, videoDB)
		}
	}
	return videoList, nil
}

func (req *CommentPublishRequest) CommentPublish(claim *utils.Claims) error {
	var videoid, commentid uint64
	if req.CommentID != "" && req.VideoID != "" {
		videoid, _ = strconv.ParseUint(req.VideoID, 10, 64)
		commentid, _ = strconv.ParseUint(req.CommentID, 10, 64)
		commentDB := &dao.CommentDB{
			UserId:     claim.ID,
			VideoId:    uint(videoid),
			ParentId:   uint(commentid),
			ChildCount: 0,
			Content:    req.Content,
		}
		err := dao.DB.Model(&dao.CommentDB{}).Create(commentDB).Error
		if err != nil {
			return err
		}
		return nil
	} else if req.CommentID != "" && req.VideoID == "" {
		commentid, _ = strconv.ParseUint(req.CommentID, 10, 64)
		commentDB := &dao.CommentDB{
			UserId:     claim.ID,
			VideoId:    0,
			ParentId:   uint(commentid),
			ChildCount: 0,
			Content:    req.Content,
		}
		err := dao.DB.Model(&dao.CommentDB{}).Create(commentDB).Error
		if err != nil {
			return err
		}
		return nil
	} else if req.VideoID != "" && req.CommentID == "" {
		videoid, _ = strconv.ParseUint(req.VideoID, 10, 64)
		commentDB := &dao.CommentDB{
			UserId:     claim.ID,
			VideoId:    uint(videoid),
			ParentId:   0,
			ChildCount: 0,
			Content:    req.Content,
		}
		err := dao.DB.Model(&dao.CommentDB{}).Create(commentDB).Error
		if err != nil {
			return err
		}
		return nil
	} else if req.CommentID == "" && req.VideoID == "" {
		return errors.New("获取信息不足")
	}
	return nil
}

func (req *CommentListRequest) CommentList() ([]dao.CommentDB, error) {
	var comments []dao.CommentDB
	if req.CommentID == "" && req.VideoID == "" {
		return comments, errors.New("查询参数不足")
	} else if (req.CommentID != "" && req.VideoID == "") || (req.CommentID != "" && req.VideoID != "") {
		parentid, _ := strconv.ParseUint(req.CommentID, 10, 64)
		var parentcomment dao.CommentDB
		err := dao.DB.Model(&dao.CommentDB{}).Where("comment_id=?", uint(parentid)).Find(&parentcomment).Error
		if gorm.IsRecordNotFoundError(err) {
			return comments, errors.New("评论不存在")
		} else if !gorm.IsRecordNotFoundError(err) {
			return comments, err
		}
		err = dao.DB.Model(&dao.CommentDB{}).Where("parent_id=?", parentcomment.ParentId).Find(&comments).Error
		if gorm.IsRecordNotFoundError(err) {
			return comments, errors.New("该评论没有子评论")
		} else if !gorm.IsRecordNotFoundError(err) {
			return comments, err
		}
		var commentlist []dao.CommentDB
		start := int(req.PageNum * req.PageSize)
		for i, comment := range comments {
			if i >= start && i < start+int(req.PageSize) {
				commentlist = append(commentlist, comment)
			}
		}
		return commentlist, nil
	} else {
		videoid, _ := strconv.ParseUint(req.VideoID, 10, 64)
		err := dao.DB.Model(&dao.CommentDB{}).Where("video_id=?", videoid).Find(&comments).Error
		if gorm.IsRecordNotFoundError(err) {
			return comments, errors.New("该评论没有子评论")
		} else if !gorm.IsRecordNotFoundError(err) {
			return comments, err
		}
		var commentlist []dao.CommentDB
		start := int(req.PageNum * req.PageSize)
		for i, comment := range comments {
			if i >= start && i < start+int(req.PageSize) {
				commentlist = append(commentlist, comment)
			}
		}
		return commentlist, nil
	}
}

func (req *CommentDeleteRequest) DeleteComment(commentid uint, userid uint) error {
	if (req.CommentID != "" && req.VideoID != "") || (req.CommentID != "" && req.VideoID == "") {
		err := dao.DB.Where("id=?", commentid).Where("user_id=?", userid).Delete(&dao.CommentDB{}).Error
		if err != nil {
			return err
		}
		return nil
	} else if req.VideoID != "" && req.CommentID == "" {
		return errors.New("未指定评论id")
	} else {
		return errors.New("参数不足")
	}
}
