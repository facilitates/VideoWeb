package dao

import (
	"github.com/jinzhu/gorm"
)

type UserDB struct {
	gorm.Model
	UserName  string
	Password  string
	AvatarUrl string
	IsBindMfa int
	MfaSecret string
}

type RelationDB struct {
	gorm.Model
	FollowerID  uint
	FollowingID uint
}

type VideoDB struct {
	gorm.Model
	UserId       uint
	VideoUrl     string
	CoverUrl     string
	Title        string
	Description  string
	VisitCount   int
	LikeCount    int
	CommentCount int
}

type CommentDB struct {
	gorm.Model
	UserId     uint
	VideoId    uint
	ParentId   uint
	LikeCount  int
	ChildCount int
	Content    string
}

type ChatDB struct {
	Content    string
	SenderId   uint
	ReceiverId uint
	Time       string
}

func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&UserDB{}).
		AutoMigrate(&VideoDB{}).
		AutoMigrate(&CommentDB{}).
		AutoMigrate(&RelationDB{}).
		AutoMigrate(&ChatDB{})
	DB.Model(&VideoDB{}).AddForeignKey("user_id", "user_db(id)", "CASCADE", "CASCADE") //User是父表
	DB.Model(&CommentDB{}).AddForeignKey("video_id", "video_db(id)", "CASCADE", "CASCADE")
}
