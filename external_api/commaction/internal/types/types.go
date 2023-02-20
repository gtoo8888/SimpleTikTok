// Code generated by goctl. DO NOT EDIT.
package types

type FavoriteActionHandlerRequest struct {
	Token      string `form:"token"`
	VideoId    int64  `form:"video_id"`
	ActionType int32  `form:"action_type"`
}

type FavoriteActionHandlerResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FavoriteListRegisterHandlerRequest struct {
	UserID int64  `form:"user_id"`
	Token  string `form:"token"`
}

type FavoriteListRegisterHandlerResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}

type Comment struct {
	Id         int64  `json:"id"          form:"id"         bson:"_id"`
	VideoId    int64  `json:"video_id"    form:"video_id"    bson:"video_id"` //视频id
	User       User   `json:"user"        form:"user"        bson:"user"`
	Content    string `json:"content"     form:"content"     bson:"content"`
	CreateDate string `json:"create_date" form:"create_date" bson:"create_date"`
}

type CommentResp struct {
	Id         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommmentActionHandlerRequest struct {
	Token       string `form:"token"`
	VideoId     int64  `form:"video_id"`
	ActionType  int32  `form:"action_type"`
	CommentText string `form:"comment_text,optional"`
	CommentId   int64  `form:"comment_id,optional"`
}

type CommmentActionHandlerResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Comment    CommentResp `json:"comment"`
}

type CommmentListHandlerRequest struct {
	Token   string `form:"token"`
	VideoId int64  `form:"video_id"`
}

type CommmentListHandlerResponse struct {
	StatusCode  int32         `json:"status_code"`
	StatusMsg   string        `json:"status_msg"`
	CommentList []CommentResp `json:"comment_list"`
}

type Video struct {
	Id            int64  `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavotite    bool   `json:"is_favorite"`
	VideoTitle    string `json:"title"`
}

type VideoTest struct {
	Id            int64  `form:"id" json:"id"`
	Author        User   `form:"author" json:"author"`
	PlayUrl       string `form:"play_url" json:"play_url"`
	CoverUrl      string `form:"cover_url" json:"cover_url"`
	FavoriteCount int64  `form:"favorite_count" json:"favorite_count"`
	CommentCount  int64  `form:"comment_count" json:"comment_count"`
	IsFavotite    bool   `form:"is_favorite" json:"is_favorite"`
}

type User struct {
	UserId        int64  `gorm:"column:user_id"        json:"id"               form:"user_id"        bson:"user_id"`
	Name          string `gorm:"column:user_nick_name" json:"name"             form:"name"           bson:"name"`
	FollowCount   int64  `gorm:"column:follow_count"   json:"follow_count"     form:"follow_count"   bson:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"   form:"follower_count" bson:"follower_count"`
	IsFollow      bool   `json:"is_follow"             form:"is_follow"        bson:"is_follow"`
}
