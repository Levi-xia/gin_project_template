package model

import (
	"project/global"
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
	err := global.App.DB.Get(&comment, "SELECT * FROM comment WHERE commentId= ?", commentId)

	return comment, err
}

func AddComment(comment Comment) (int64, error) {

	sqlStr := "INSERT INTO comment (content, createTime, updateTime) VALUES (:content, :createTime, :updateTime)"

	result, err := global.App.DB.NamedExec(sqlStr, comment)

	if err != nil {
		return 0, err
	}
	newCommentId, err := result.LastInsertId()

	return newCommentId, nil
}
