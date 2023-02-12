package dto

import "project/common"

type GetCommentReq struct {
	CommentId int64 `form:"commentId" json:"commentId" validate:"required,gt=0"`
}
type GetCommentResp struct {
	CommentId int64  `json:"commentId"`
	Content   string `json:"content"`
}

func (req GetCommentReq) GetMessages() common.ValidatorMessages {
	return common.ValidatorMessages{
		"CommentId.required": "评论Id不能为空",
		"CommentId.gt":       "评论Id需要大于0",
	}
}

type AddCommentReq struct {
	UserId     int64  `form:"userId" json:"userId" validate:"required,gt=0"`
	TopicType  string `form:"topicType" json:"topicType" validate:"required,oneof=blog look"`
	TopicId    int64  `form:"topicId" json:"topicId" validate:"required,gt=0"`
	ReplyId    int64  `form:"replyId" json:"replyId"`
	Content    string `form:"content" json:"content" validate:"required"`
	TextStruct string `form:"textStruct" json:"textStruct"`
	Owner      uint   `form:"owner" json:"owner" validate:"oneof=0 1"`
}
type AddCommentResp struct {
	CommentId int64 `json:"commentId"`
}

func (req AddCommentReq) GetMessages() common.ValidatorMessages {
	return common.ValidatorMessages{
		"UserId.required":    "用户Id不能为空",
		"UserId.gt":          "用户Id需要大于0",
		"TopicType.required": "评论类型不能为空",
		"TopicType.oneof":    "评论类型非法",
		"TopicId.required":   "评论对象Id不能为空",
		"TopicId.gt":         "评论对象Id需要大于0",
		"Content.required":   "评论内容不能为空",
		"Owner.oneof":        "评论拥有人参数不合法",
	}
}
