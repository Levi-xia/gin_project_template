package dto

import "project/common"

type GetCommentReq struct {
	CommentId int64 `form:"commentId" json:"commentId" validate:"required,gt=0"`
}

func (req GetCommentReq) GetMessages() common.ValidatorMessages {
	return common.ValidatorMessages{
		"CommentId.required": "评论Id不能为空",
		"CommentId.gt":       "评论Id需要大于0",
	}
}

type GetCommentResp struct {
	CommentId int64
}
