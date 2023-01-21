package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/mongodb"
	"SimpleTikTok/oprations/sql"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type CommmentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommmentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommmentActionLogic {
	return &CommmentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommmentActionLogic) CommmentAction(req *types.CommmentActionHandlerRequest) (resp *types.CommmentActionHandlerResponse, err error) {
	// todo: add your logic here and delete this line
	//parse token
	flag, userId := tools.CheckToke(req.Token)
	if flag == false {
		return nil, errors.New("parse Token failed")
	}

	//get collection from mongodb
	mongoUser := "admin"
	mongoPwd := "admin"
	mongoUrl := "192.168.31.132:27017"
	url := fmt.Sprintf("mongodb://%v:%v@%v", mongoUser, mongoPwd, mongoUrl)
	collection, err := mongodb.Connect("tiktok", "comment", url)
	if err != nil {
		return nil, err
	}

	actionType := req.ActionType
	//delete comment
	if actionType == 2 {
		commentId := req.CommentId
		filter := bson.D{{
			Key:   "_id",
			Value: commentId,
		}}
		_, err = collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return nil, err
		}
		resp.StatusCode = 0
		resp.StatusMsg = fmt.Sprintf("delete success")
	} else {
		db, err := sql.SqlConnect()
		if err != nil {
			return nil, err
		}

		var user types.User
		err = db.Table("user_info").Where("user_id=?", userId).First(&user).Error
		if err != nil {
			return nil, err
		}
		videoId := req.VideoId
		content := req.CommentText
		date := time.Now()
		createDate := fmt.Sprintf("%d-%v", date.Month(), date.Day())
		comment := types.Comment{
			VideoId:    videoId,
			User:   	user,
			Content:    content,
			CreateDate: createDate,
		}
		_, err = collection.InsertOne(context.Background(), comment)
		if err != nil {
			return nil, err
		}
		resp.StatusCode = 0
		resp.StatusMsg = "insert success"
		resp.Comment = comment
	}
	return
}
