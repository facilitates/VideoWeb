package videoweb

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
	"videoweb/dao"
	"videoweb/pkg/utils"
)

func (req *VideoFeedRequest) VideoFeed(latestTime time.Time) ([]dao.VideoDB, error) {
	var videoDB []dao.VideoDB
	if req.LatestTime == "" {
		err := dao.DB.Model(&dao.VideoDB{}).Find(&videoDB).Error
		if gorm.IsRecordNotFoundError(err) {
			return videoDB, errors.New("视频不存在")
		}
		return videoDB, err
	} else {
		err := dao.DB.Model(&dao.VideoDB{}).Where("created_at>=?", latestTime).Find(&videoDB).Error
		if gorm.IsRecordNotFoundError(err) {
			return videoDB, errors.New("视频不存在")
		} else if !gorm.IsRecordNotFoundError(err) {
			return videoDB, err
		}
		return videoDB, nil
	}
}

func (req *UploadVideoRequest) UploadVideo(claim *utils.Claims, filepath string) error {
	var videoDB dao.VideoDB
	videoDB = dao.VideoDB{
		UserId:      claim.ID,
		VideoUrl:    filepath,
		Title:       req.Title,
		Description: req.Description,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err := dao.DB.Model(&dao.VideoDB{}).Save(&videoDB).Error
	return err
}

func (req *VideoListRequest) VideoList(userid uint) ([]dao.VideoDB, int, error) {
	var videos []dao.VideoDB
	var count int
	err := dao.DB.Model(&dao.VideoDB{}).Where("user_id=?", userid).Find(&videos).Count(&count).Error
	if gorm.IsRecordNotFoundError(err) {
		return videos, count, errors.New("用户未上传视频")
	} else if !gorm.IsRecordNotFoundError(err) {
		return videos, count, err
	}
	return videos, count, nil
}

func (req *PopularVideoRequest) PopularVideo() ([]dao.VideoDB, error) {
	key := "videolike"
	var videos []dao.VideoDB
	ranklist := dao.Redisdb.ZRevRangeWithScores(key, 0, -1)
	VideoRankList, err := ranklist.Result()
	if err != nil {
		return videos, err
	}
	var videorank []dao.VideoDB
	var start = int(req.PageSize * req.PageNum)
	for i, z := range VideoRankList {
		if i+1 >= start && i+1 <= start+int(req.PageSize) {
			var video dao.VideoDB
			member := z.Member.(string)
			likecount := z.Score
			dao.DB.First(&video, member)
			video.LikeCount = int(likecount)
			videorank = append(videorank, video)
		} else {
			continue
		}
	}
	return videorank, nil
}
