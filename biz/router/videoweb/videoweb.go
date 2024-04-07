// Code generated by hertz generator. DO NOT EDIT.

package videoweb

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	videoweb "videoweb/biz/handler/videoweb"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_auth := root.Group("/auth", _authMw()...)
		{
			_mfa := _auth.Group("/mfa", _mfaMw()...)
			_mfa.POST("/bind", append(_mfabindMw(), videoweb.MfaBind)...)
			_mfa.GET("/qrcode", append(_mfaqrcodeMw(), videoweb.MfaQrcode)...)
		}
	}
	{
		_comment := root.Group("/comment", _commentMw()...)
		_comment.DELETE("/delete", append(_commentdeleteMw(), videoweb.CommentDelete)...)
		_comment.GET("/list", append(_commentlistMw(), videoweb.CommentList)...)
		_comment.POST("/publish", append(_commentpublishMw(), videoweb.CommentPublish)...)
	}
	{
		_follower := root.Group("/follower", _followerMw()...)
		_follower.GET("/list", append(_followerlistMw(), videoweb.FollowerList)...)
	}
	{
		_following := root.Group("/following", _followingMw()...)
		_following.GET("/list", append(_followinglistMw(), videoweb.FollowingList)...)
	}
	{
		_friends := root.Group("/friends", _friendsMw()...)
		_friends.GET("/list", append(_friendslistMw(), videoweb.FriendsList)...)
	}
	{
		_like := root.Group("/like", _likeMw()...)
		_like.POST("/action", append(_likeactionMw(), videoweb.LikeAction)...)
		_like.GET("/list", append(_likelistMw(), videoweb.LikeList)...)
	}
	{
		_relation := root.Group("/relation", _relationMw()...)
		_relation.POST("/action", append(_relationactionMw(), videoweb.RelationAction)...)
	}
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/info", append(_infoMw(), videoweb.Info)...)
		_user.POST("/login", append(_loginMw(), videoweb.Login)...)
		_user.POST("/register", append(_registerMw(), videoweb.Register)...)
		{
			_avatar := _user.Group("/avatar", _avatarMw()...)
			_avatar.PUT("/upload", append(_avataruploadMw(), videoweb.AvatarUpload)...)
		}
	}
	{
		_video := root.Group("/video", _videoMw()...)
		_video.GET("/feed", append(_videofeedMw(), videoweb.VideoFeed)...)
		_video.GET("/list", append(_videolistMw(), videoweb.VideoList)...)
		_video.GET("/popular", append(_popularvideoMw(), videoweb.PopularVideo)...)
		_video.POST("/publish", append(_uploadvideoMw(), videoweb.UploadVideo)...)
		_video.POST("/search", append(_searchvideoMw(), videoweb.SearchVideo)...)
	}
}