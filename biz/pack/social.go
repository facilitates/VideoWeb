package pack

import (
	"strconv"
	"videoweb/biz/model/videoweb"
	"videoweb/dao"
)

func BuildFollowingList(followings []dao.UserDB, count int, state int) videoweb.FollowingListResponse {
	followinglist := BuildFollowings(followings)
	if state == 1 {
		return videoweb.FollowingListResponse{
			Base: &videoweb.BaseResp{
				Code: 200,
				Msg:  "查询用户关注列表成功",
			},
			Data: &videoweb.Data{
				Items: followinglist,
				Total: int64(count),
			},
		}
	} else {
		return videoweb.FollowingListResponse{
			Base: &videoweb.BaseResp{
				Code: 200,
				Msg:  "查询用户粉丝列表成功",
			},
			Data: &videoweb.Data{
				Items: followinglist,
				Total: int64(count),
			},
		}
	}
}

func BuildFollowings(followings []dao.UserDB) []*videoweb.Following {
	var followinglist []*videoweb.Following
	for _, following := range followings {
		result := BuildFollowing(following)
		followinglist = append(followinglist, result)
	}
	return followinglist
}

func BuildFollowing(following dao.UserDB) *videoweb.Following {
	id := strconv.FormatUint(uint64(following.ID), 10)
	return &videoweb.Following{
		ID:        id,
		UserName:  following.UserName,
		AvatarURL: following.AvatarUrl,
	}
}

func BuildRelationActionResp(state int) videoweb.RelationActionResponse {
	if state == 1 {
		return videoweb.RelationActionResponse{
			Base: &videoweb.BaseResp{
				Code: 200,
				Msg:  "取消关注成功",
			},
		}
	} else {
		return videoweb.RelationActionResponse{
			Base: &videoweb.BaseResp{
				Code: 200,
				Msg:  "关注成功",
			},
		}
	}
}

func BuildFriendsListResp(onlineFriends []dao.UserDB, count int) videoweb.FriendsListResponse {
	list := BuildFollowings(onlineFriends)
	return videoweb.FriendsListResponse{
		Base: &videoweb.BaseResp{
			Code: 200,
			Msg:  "查询用户关注列表成功",
		},
		Data: &videoweb.Data{
			Items: list,
			Total: int64(count),
		},
	}
}
