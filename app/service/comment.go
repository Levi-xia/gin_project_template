package service

import (
	"github.com/gin-gonic/gin"
	"project/app/model"
)

func GetCommentById(c *gin.Context, commentId int64) (model.Comment, error) {

	comment, err := model.GetCommentById(commentId)
	return comment, err
}
