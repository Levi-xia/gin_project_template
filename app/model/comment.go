package model

import (
	"project/app/model/gdao"
)

type Comment struct {
	CommentId  int64  `db:"commentId" json:"commentId"`
	UserId     int64  `db:"userId" json:"userId"`
	TopicType  string `db:"topicType" json:"topicType"`
	TopicId    int64  `db:"topicId" json:"topicId"`
	ReplyId    int64  `db:"replyId" json:"replyId"`
	Content    string `db:"content" json:"content"`
	TextStruct string `db:"textStruct" json:"textStruct"`
	Owner      uint   `db:"owner" json:"owner"`
	Stat       uint   `db:"stat" json:"stat"`
	CreateTime int64  `db:"createTime" json:"createTime"`
	UpdateTime int64  `db:"updateTime" json:"updateTime"`
}

func GetCommentById(commentId int64) (Comment, error) {
	comment := Comment{}
	point := gdao.GetEndPoint[Comment]{
		Model: &comment,
		Table: "comment",
		Conditions: map[string]any{
			"commentId=": commentId,
		},
		Fields: []string{"*"},
	}
	err := point.Get()
	return comment, err
}

func AddComment(userId int64, topicId int64, content string, topicType string, replyId int64, owner uint) (int64, error) {
	point := gdao.InsertEndpoint[Comment]{
		Table: "comment",
		Rows: map[string]any{
			"topicId":   topicId,
			"topicType": topicType,
			"userId":    userId,
			"replyId":   replyId,
			"content":   content,
			"owner":     owner,
		},
	}
	insertId, err := point.Insert()
	return insertId, err

}
