package service

import (
	"github.com/gin-gonic/gin"
	"project/app/model"
)

func GetCommentById(c *gin.Context, commentId int64) (model.Comment, error) {
	comment, err := model.GetCommentById(commentId)
	return comment, err
}

func AddComment(c *gin.Context, comment model.Comment) (int64, error) {
	userId := comment.UserId
	topicId := comment.TopicId
	content := comment.Content
	topicType := comment.TopicType
	replyId := comment.ReplyId
	owner := comment.Owner

	id, err := model.AddComment(userId, topicId, content, topicType, replyId, owner)

	return id, err
}
