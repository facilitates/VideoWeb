package videoweb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"time"
	"videoweb/dao"
)

func (req *RelationActionRequest) RelationAction(followerid uint, followingid uint) (int, error) {
	var count int
	dao.DB.Model(&dao.RelationDB{}).Where("follower_id=?", followerid).Where("following_id=?", followingid).Count(&count)
	if req.ActionType == 0 {
		if count == 1 {
			return 2, errors.New("已经关注该用户")
		} else if count == 0 {
			relationDB := &dao.RelationDB{
				FollowerID:  followerid,
				FollowingID: followingid,
			}
			err := dao.DB.Model(&dao.RelationDB{}).Create(&relationDB).Error
			if err != nil {
				return 2, err
			}
			return 0, nil
		}
	} else if req.ActionType == 1 {
		if count == 0 {
			return 2, errors.New("未关注该用户")
		} else if count == 1 {
			err := dao.DB.Where("follower_id=?", followerid).Where("following_id=?", followingid).Delete(&dao.RelationDB{}).Error
			if err != nil {
				return 2, err
			}
			return 1, nil
		}
	} else {
		return 2, errors.New("未指定操作类型")
	}
	return 2, errors.New("错误")
}

func (req *FollowingListRequest) FollowingList(followerid uint) ([]dao.UserDB, int, error) {
	var relations []dao.RelationDB
	var followings []dao.UserDB
	var count int
	err := dao.DB.Model(&dao.RelationDB{}).Where("follower_id=?", followerid).Find(&relations).Count(&count).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			log.Println(err)
			return followings, count, errors.New("用户没有任何关注")
		} else if !gorm.IsRecordNotFoundError(err) {
			log.Println(err)
			return followings, count, err
		}
	}
	for _, relation := range relations {
		var user dao.UserDB
		followingid := strconv.FormatUint(uint64(relation.FollowingID), 10)
		err := dao.DB.Model(&dao.UserDB{}).Where("id=?", followingid).First(&user).Error
		if err != nil {
			return followings, count, err
		}
		followings = append(followings, user)
	}
	return followings, count, nil
}

func (req *FollowerListRequest) FollowerList(userid uint) ([]dao.UserDB, int, error) {
	var followers []dao.RelationDB
	var users []dao.UserDB
	var count int
	err := dao.DB.Model(&dao.RelationDB{}).Where("following_id=?", userid).Find(&followers).Count(&count).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return users, count, errors.New("用户没有粉丝")
		} else {
			return users, count, err
		}
	}
	for _, follower := range followers {
		var user dao.UserDB
		err := dao.DB.Model(&dao.UserDB{}).Where("id=?", follower.FollowerID).Find(&user).Error
		if err != nil {
			return users, count, err
		}
		users = append(users, user)
	}
	return users, count, nil
}

func (req *FriendsListRequest) FriendList(userid uint) ([]dao.UserDB, error) {
	var friends []dao.UserDB
	var count int
	err := dao.DB.Model(&dao.UserDB{}).Where("follower_id=?", userid).Find(&friends).Count(&count).Error
	if err != nil {
		return friends, err
	} else {
		return friends, nil
	}
}

func SendMessageToUser(content dao.ChatDB, userkey string) error {
	contentBytes, err := json.Marshal(content)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(time.Now())
	err = dao.Redisdb.ZAdd(userkey, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: contentBytes,
	}).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChatHistoryWithUser(userid uint) ([]dao.ChatDB, error) {
	var chathistory []dao.ChatDB
	err := dao.DB.Model(&dao.ChatDB{}).Where("receiver_id=?", userid).Find(&chathistory).Error
	if err != nil {
		return chathistory, err
	}
	return chathistory, nil
}

func SearchUnrecHistory(chatkey string) ([]dao.ChatDB, error) {
	var chathistory []dao.ChatDB
	UnrecHistory, err := dao.Redisdb.ZRangeWithScores(chatkey, 0, -1).Result()
	if err != nil {
		return chathistory, err
	}
	for _, i := range UnrecHistory {
		chat := dao.ChatDB{
			Content: i.Member.(string),
			Time:    strconv.FormatFloat(i.Score, 'f', -1, 64),
		}
		chathistory = append(chathistory, chat)
	}
	err = dao.Redisdb.Del(chatkey).Err()
	if err != nil {
		return chathistory, err
	}
	return chathistory, nil
}

func SendMessageToGroup(groupkey string, content string, senderid uint) error {
	Content := dao.ChatDB{
		Content:  content,
		SenderId: senderid,
	}
	contentBytes, err := json.Marshal(Content)
	if err != nil {
		return err
	}
	err = dao.Redisdb.ZAddNX(groupkey, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: contentBytes,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

func SaveUnrecHistory(packedUnrec []dao.ChatDB) error {
	for _, a := range packedUnrec {
		err := dao.DB.Model(&dao.ChatDB{}).Save(a).Error
		return err
	}
	return nil
}
